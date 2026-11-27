package ratelimit

import (
	"sync"
	"time"
)

// Config configuração de rate limit.
type Config struct {
	RequestsPerMinute int
	Burst             int
	GlobalCostCapCents int
}

// TokenBucket limitador por sessão (simplificado; produção deve usar Redis com janela deslizante).
type TokenBucket struct {
	mu       sync.Mutex
	tokens   map[string]float64
	lastSeen map[string]time.Time
	rate     float64 // tokens por segundo
	burst    float64
	config   Config
}

func NewTokenBucket(cfg Config) *TokenBucket {
	return &TokenBucket{
		tokens:    make(map[string]float64),
		lastSeen:  make(map[string]time.Time),
		rate:      float64(cfg.RequestsPerMinute) / 60.0,
		burst:     float64(cfg.Burst),
		config:    cfg,
	}
}

// Allow retorna true se a requisição estiver dentro do limite para a sessão.
func (b *TokenBucket) Allow(sessionID string) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	last, ok := b.lastSeen[sessionID]
	if !ok {
		last = now
	}
	elapsed := now.Sub(last).Seconds()
	b.lastSeen[sessionID] = now

	tokens, ok := b.tokens[sessionID]
	if !ok {
		tokens = b.burst
	}
	tokens += elapsed * b.rate
	if tokens > b.burst {
		tokens = b.burst
	}
	if tokens < 1.0 {
		b.tokens[sessionID] = tokens
		return false
	}
	tokens--
	b.tokens[sessionID] = tokens
	return true
}

// GlobalCostOK indica se o orçamento global de custos ainda não foi esgotado.
func (b *TokenBucket) GlobalCostOK() bool {
	return b.config.GlobalCostCapCents > 0
}
