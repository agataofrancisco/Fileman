package main

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ao.libre.fileman/backend/jev-gateway/internal/auth"
	"ao.libre.fileman/backend/jev-gateway/internal/handlers"
	"ao.libre.fileman/backend/jev-gateway/internal/models"
	"ao.libre.fileman/backend/jev-gateway/internal/ratelimit"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	jwtSecret := os.Getenv("JEV_GATEWAY_JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-only-insecure-secret-change-me"
		logger.Warn("usando JWT secret de desenvolvimento; defina JEV_GATEWAY_JWT_SECRET em produção")
	}

	jevAPIKey := os.Getenv("JEV_API_KEY")
	if jevAPIKey == "" {
		logger.Error("JEV_API_KEY não definida; o gateway não funcionará sem a chave do TypeSafe/Jev")
		os.Exit(1)
	}

	store := auth.NewMemoryStore(jwtSecret)
	limiter := ratelimit.NewTokenBucket(rateLimitConfig())
	svc := handlers.NewGatewayService(jevAPIKey, store, limiter, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/v1/session", svc.CreateSession)
	mux.HandleFunc("/v1/classify", svc.Classify)

	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	server := &http.Server{
		Addr:              addr,
		Handler:           requestLogger(logger)(mux),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	go func() {
		logger.Info("gateway Jev escutando", "addr", addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("erro no servidor HTTP", "err", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("desligando gateway...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("erro ao desligar", "err", err)
	}
	logger.Info("gateway desligado")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.HealthResponse{Status: "ok"})
}

func requestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Info("requisição",
				"method", r.Method,
				"path", r.URL.Path,
				"duration_ms", time.Since(start).Milliseconds(),
			)
			// Nunca loga query params, headers de autorização ou corpo da requisição/resposta.
		})
	}
}

func rateLimitConfig() ratelimit.Config {
	return ratelimit.Config{
		RequestsPerMinute: 60,
		Burst:             10,
		GlobalCostCapCents: 5000, // trava global de custos (exemplo, configurável)
	}
}
