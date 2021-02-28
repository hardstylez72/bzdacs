package session

import (
	"github.com/google/uuid"
	"time"
)

type Service struct {
	store *store
}

func NewSessionService() *Service {
	return &Service{NewSessionStore()}
}

type Session struct {
	ExpiresAt time.Time
	Login     string
	Password  string
	Token     string
}

func (s *Service) GenerateSession(login, password string, expiresAt time.Time) *Session {
	return &Session{
		ExpiresAt: expiresAt,
		Login:     login,
		Password:  password,
		Token:     uuid.New().String(),
	}
}

func (s *Service) GetByToken(token string) (*Session, error) {
	session, err := s.store.GetByToken(token)
	if err != nil {
		return nil, err
	}
	if session.ExpiresAt.Before(time.Now()) {
		_ = s.store.Delete(session.Token)
		return nil, ErrInvalidSession
	}

	return session, nil
}

func (s *Service) Set(session *Session) (string, error) {
	if session.Token == "" || session.Password == "" || session.Login == "" || session.ExpiresAt.IsZero() {
		return "", ErrInvalidSession
	}
	token, err := s.store.Set(session)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Clear(token string) error {
	return s.store.Delete(token)
}
