package kvstore

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentAccess(t *testing.T) {
	s := NewKVStore()
	var wg sync.WaitGroup
	numGoroutines := 100
	wg.Add(numGoroutines)

	for i := range numGoroutines {
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)

			// Concurrently set a value.
			s.Set(key, value)

			// Concurrently get a value.
			_, _ = s.Get(key)
		}(i)
	}

	wg.Wait()
	for i := range numGoroutines {
		key := fmt.Sprintf("key%d", i)
		expectedValue := fmt.Sprintf("value%d", i)

		val, err := s.Get(key)
		if err != nil {
			t.Errorf("Expected key %s to exist, but got error: %v", key, err)
		}
		if val != expectedValue {
			t.Errorf("Expected value for key %s to be %s, but got %s", key, expectedValue, val)
		}
	}
}

func TestStore(t *testing.T) {
	s := NewKVStore()

	if s.Size() != 0 {
		t.Errorf("store is not empty")
	}

	value, error := s.Get("key")
	if error != ErrNotFound || value != "" {
		t.Errorf("store has extra keys")
	}

	value, error = s.Delete("key")
	if error != ErrNotFound || value != "" {
		t.Errorf("store has extra keys")
	}

	s.Set("key", "value")
	if s.Size() != 1 {
		t.Errorf("store has extra keys")
	}

	value, error = s.Get("key")
	if error != nil || value != "value" {
		t.Errorf("store didnt set key")
	}

	value, error = s.Delete("key")
	if error != nil || value != "value" {
		t.Errorf("store didnt set key")
	}

	if s.Size() != 0 {
		t.Errorf("store is not empty")
	}
}
