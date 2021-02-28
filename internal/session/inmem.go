package session

import (
	"errors"
	"sync"
)

type inmem struct {
	m *sync.Map
}

var ErrNotFound = errors.New("not found in store")

func NewInMemStorage() *inmem {
	return &inmem{m: &sync.Map{}}
}

func (s *inmem) Set(key, value interface{}) {
	s.m.Store(key, value)
}

func (s *inmem) Get(key interface{}) (interface{}, error) {
	val, exist := s.m.Load(key)
	if !exist {
		return nil, ErrNotFound
	}
	return val, nil
}

func (s *inmem) Delete(key interface{}) {
	s.m.Delete(key)
}
