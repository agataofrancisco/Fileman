package auth

import (
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Session sessão anónima de utilizador (sem conta obrigatória no MVP).
type Session struct {
	ID        string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// Store interface de armazenamento de sessões (memória por enquanto; substituir por Redis/DynamoDB em produção).
type Store interface {
	CreateSession(ttl time.Duration) (*Session, error)
	GetSession(id string) (*Session, bool)
	DeleteSession(id string)
}

// MemoryStore implementação em memória da Store.
type MemoryStore struct {
	mu      sync.RWMutex
	sessions map[string]*Session
	secret  []byte
	logger  *slog.Logger
}

func NewMemoryStore(jwtSecret string) *MemoryStore {
	return &MemoryStore{
		sessions: make(map[string]*Session),
		secret:   []byte(jwtSecret),
		logger:   slog.Default(),
	}
}

func (s *MemoryStore) CreateSession(ttl time.Duration) (*Session, error) {
	id := uuid.NewString()
	now := time.Now()
	sess := &Session{
		ID:        id,
		CreatedAt: now,
		ExpiresAt: now.Add(ttl),
	}
	s.mu.Lock()
	s.sessions[id] = sess
	s.mu.Unlock()
	return sess, nil
}

func (s *MemoryStore) GetSession(id string) (*Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sess, ok := s.sessions[id]
	if !ok || time.Now().After(sess.ExpiresAt) {
		return nil, false
	}
	return sess, true
}

func (s *MemoryStore) DeleteSession(id string) {
	s.mu.Lock()
	delete(s.sessions, id)
	s.mu.Unlock()
}

var ErrInvalidToken = errors.New("token inválido")

// ValidateJWT valida um JWT de sessão (não guarda dados pessoais no payload).
func ValidateJWT(tokenString string, secret []byte) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return secret, nil
	})
	if err != nil || !token.Valid {
		return "", ErrInvalidToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrInvalidToken
	}
	sessionID, ok := claims["sid"].(string)
	if !ok || sessionID == "" {
		return "", ErrInvalidToken
	}
	return sessionID, nil
}
