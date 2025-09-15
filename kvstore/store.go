package kvstore

import (
	"fmt"
	"sync"
)

var ErrNotFound = fmt.Errorf("key not found")

// KVStore is a concurrently safe, in-memory key-value store.
type KVStore struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewKVStore creates and initializes a new KVStore.
func NewKVStore() *KVStore {
	return &KVStore{
		store: make(map[string]string),
	}
}

func (s *KVStore) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.store)
}

func (s *KVStore) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
}

func (s *KVStore) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if value, exists := s.store[key]; exists {
		return value, nil
	}

	return "", ErrNotFound
}

func (s *KVStore) Delete(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if value, exists := s.store[key]; exists {
		delete(s.store, key)
		return value, nil
	}
	return "", ErrNotFound
}
