package session

import (
	"errors"
)

type store struct {
	inmem *inmem
}

var ErrInvalidSession = errors.New("invalid session")

func NewSessionStore() *store {
	return &store{NewInMemStorage()}
}

func (s *store) GetByToken(token string) (*Session, error) {
	val, err := s.inmem.Get(token)
	if err != nil {
		return nil, err
	}
	session, ok := val.(*Session)
	if !ok {
		return nil, ErrInvalidSession
	}

	return session, nil
}

func (s *store) Set(session *Session) (string, error) {
	if session.Token == "" || session.Password == "" || session.Login == "" || session.ExpiresAt.IsZero() {
		return "", ErrInvalidSession
	}
	s.inmem.Set(session.Token, session)

	return session.Token, nil
}

func (s *store) Delete(token string) error {
	s.inmem.Delete(token)
	return nil
}
