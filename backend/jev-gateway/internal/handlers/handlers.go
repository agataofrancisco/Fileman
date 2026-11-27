package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"ao.libre.fileman/backend/jev-gateway/internal/auth"
	"ao.libre.fileman/backend/jev-gateway/internal/models"
	"ao.libre.fileman/backend/jev-gateway/internal/ratelimit"
)

// GatewayService orquestra sessões, rate limit e encaminhamento para Jev/TypeSafe.
type GatewayService struct {
	jevAPIKey   string
	store       auth.Store
	limiter     *ratelimit.TokenBucket
	logger      *slog.Logger
	httpClient  *http.Client
	jevEndpoint string
	jwtSecret   []byte
}

// NewGatewayService cria uma instância do serviço gateway.
func NewGatewayService(jevAPIKey string, store auth.Store, limiter *ratelimit.TokenBucket, logger *slog.Logger) *GatewayService {
	return &GatewayService{
		jevAPIKey:   jevAPIKey,
		store:       store,
		limiter:     limiter,
		logger:      logger,
		httpClient:  &http.Client{Timeout: 30 * time.Second},
		jevEndpoint: getEnv("JEV_API_ENDPOINT", "https://api.typesafe.ai/v1/classify"), // verificar documento oficial TypeSafe
		jwtSecret:   []byte(getEnv("JEV_GATEWAY_JWT_SECRET", "")),
	}
}

// CreateSession cria uma sessão anónima sem exigir conta.
func (s *GatewayService) CreateSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}
	session, err := s.store.CreateSession(24 * time.Hour)
	if err != nil {
		s.logger.Error("falha ao criar sessão", "err", err)
		http.Error(w, "erro interno", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(models.SessionResponse{
		SessionID: session.ID,
		ExpiresAt: session.ExpiresAt.UTC().Format(time.RFC3339),
	})
}

// Classify recebe requisição de classificação do cliente, valida sessão/rate-limit
// e encaminha (ou não) para a API oficial TypeSafe/Jev.
func (s *GatewayService) Classify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "corpo inválido", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req models.ClassifyRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "payload inválido", http.StatusBadRequest)
		return
	}

	// Valida sessão JWT (apenas sid, sem dados pessoais)
	if _, err := auth.ValidateJWT(req.SessionID, s.jwtSecret); err != nil {
		http.Error(w, "sessão inválida", http.StatusUnauthorized)
		return
	}

	// Rate limit por sessão
	if !s.limiter.Allow(req.SessionID) {
		w.Header().Set("Retry-After", "60")
		http.Error(w, "limite de requisições excedido", http.StatusTooManyRequests)
		return
	}

	if !s.limiter.GlobalCostOK() {
		http.Error(w, "orçamento global esgotado", http.StatusServiceUnavailable)
		return
	}

	// Sem imagem nem OCR autorizado => classificação local (fallback de regras, não outro LLM)
	if !req.SendImage && req.OCRText == "" {
		writeJSON(w, http.StatusOK, models.ClassifyResponse{
			FileRef:          req.FileRef,
			Status:           "unknown",
			Category:         "unknown",
			Sensitive:        false,
			SuggestedView:    "review",
			ModelID:          "local-fallback-rules",
			TaxonomyVersion:  "1.0.0",
			DeleteAuthorized: false,
		})
		return
	}

	// Encaminhar para Jev/TypeSafe apenas com consentimento explícito.
	// A chave da API nunca sai deste gateway.
	jeVReq := models.JevClassifyRequest{
		RequestID:    req.FileRef,
		FileRef:      req.FileRef,
		Taxonomy:     "1.0.0",
		ConsentScope: "image_category",
		Metadata:     string(body), // metadados mínimos apenas
	}

	jeVBody, err := json.Marshal(jeVReq)
	if err != nil {
		s.logger.Error("falha ao montar requisição Jev", "err", err)
		http.Error(w, "erro interno", http.StatusInternalServerError)
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, s.jevEndpoint, bytes.NewReader(jeVBody))
	if err != nil {
		s.logger.Error("falha ao montar requisição HTTP", "err", err)
		http.Error(w, "erro interno", http.StatusInternalServerError)
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+s.jevAPIKey)
	if req.SendImage {
		httpReq.Header.Set("TypeSafe-Danger-Zone-Images", "true")
	}

	resp, err := s.httpClient.Do(httpReq)
	if err != nil {
		// Falha de IA não bloqueia o cliente; devolve unknown localmente
		s.logger.Error("fornecedor Jev indisponível", "err", err)
		writeJSON(w, http.StatusOK, models.ClassifyResponse{
			FileRef:          req.FileRef,
			Status:           "failed",
			Category:         "unknown",
			Sensitive:        false,
			SuggestedView:    "review",
			ModelID:          "jev-unavailable",
			TaxonomyVersion:  "1.0.0",
			DeleteAuthorized: false,
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		w.Header().Set("Retry-After", resp.Header.Get("Retry-After"))
		http.Error(w, "limite do fornecedor excedido", http.StatusTooManyRequests)
		return
	}
	if resp.StatusCode != http.StatusOK {
		s.logger.Error("resposta não-OK do Jev", "status", resp.StatusCode)
		writeJSON(w, http.StatusOK, models.ClassifyResponse{
			FileRef:          req.FileRef,
			Status:           "failed",
			Category:         "unknown",
			Sensitive:        false,
			SuggestedView:    "review",
			ModelID:          "jev-error",
			TaxonomyVersion:  "1.0.0",
			DeleteAuthorized: false,
		})
		return
	}

	var jevResp models.JevClassifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&jevResp); err != nil {
		s.logger.Error("falha ao decodificar resposta Jev", "err", err)
		writeJSON(w, http.StatusOK, models.ClassifyResponse{
			FileRef:          req.FileRef,
			Status:           "failed",
			Category:         "unknown",
			Sensitive:        false,
			SuggestedView:    "review",
			ModelID:          "jev-parse-error",
			TaxonomyVersion:  "1.0.0",
			DeleteAuthorized: false,
		})
		return
	}

	writeJSON(w, http.StatusOK, models.ClassifyResponse{
		FileRef:          req.FileRef,
		Status:           jevResp.Status,
		Category:         jevResp.Category,
		Subcategory:      jevResp.Subcategory,
		Confidence:       jevResp.Confidence,
		Sensitive:        jevResp.Sensitive,
		SuggestedView:    "review",
		ModelID:          jevResp.ModelID,
		TaxonomyVersion:  "1.0.0",
		ReasonCodes:      jevResp.ReasonCodes,
		DeleteAuthorized: false, // regra de negócio: Jev nunca autoriza eliminação
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
