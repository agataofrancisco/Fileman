package models

// HealthResponse resposta do endpoint de saúde.
type HealthResponse struct {
	Status string `json:"status"`
}

// SessionResponse resposta da criação de sessão anónima.
type SessionResponse struct {
	SessionID string `json:"sessionId"`
	ExpiresAt string `json:"expiresAt"`
}

// ClassifyRequest requisição do cliente para classificação (sanitizada).
type ClassifyRequest struct {
	SessionID   string `json:"sessionId"`
	FileRef     string `json:"fileRef"`       // ID opaco local, nunca caminho real
	Mime        string `json:"mime"`
	SizeBytes   int64  `json:"sizeBytes"`
	ModifiedAt  string `json:"modifiedAt"`
	SourceKind  string `json:"sourceKind"`    // ex.: whatsapp_media
	OCRText     string `json:"ocrText,omitempty"` // apenas com consentimento
	SendImage   bool   `json:"sendImage"`      // opt-in explícito de envio de imagem
}

// ClassifyResponse resposta normalizada devolvida ao cliente.
type ClassifyResponse struct {
	FileRef          string   `json:"fileRef"`
	Status           string   `json:"status"`            // classified | unknown | failed | pending
	Category         string   `json:"category"`
	Subcategory      string   `json:"subcategory,omitempty"`
	Confidence       *float64 `json:"confidence,omitempty"`
	Sensitive        bool     `json:"sensitive"`
	SuggestedView    string   `json:"suggestedView"`     // ex.: review
	ModelID          string   `json:"modelId"`
	TaxonomyVersion  string   `json:"taxonomyVersion"`
	ReasonCodes      []string `json:"reasonCodes,omitempty"`
	DeleteAuthorized bool     `json:"deleteAuthorized"`  // sempre false
}

// JevClassifyRequest envio para a API oficial TypeSafe/Jev (preenchido pelo gateway).
type JevClassifyRequest struct {
	RequestID    string `json:"requestId"`
	FileRef      string `json:"fileRef"`
	Taxonomy     string `json:"taxonomyVersion"`
	ConsentScope string `json:"consentScope"`
	Metadata     string `json:"metadata"` // JSON string dos metadados mínimos
}

// JevClassifyResponse resposta bruta da TypeSafe/Jev (mapeada pelo gateway).
type JevClassifyResponse struct {
	Status      string   `json:"status"`
	Category    string   `json:"category"`
	Subcategory string   `json:"subcategory,omitempty"`
	Confidence  *float64 `json:"confidence,omitempty"`
	Sensitive   bool     `json:"sensitive"`
	ReasonCodes []string `json:"reasonCodes,omitempty"`
	ModelID     string   `json:"modelId"`
}
