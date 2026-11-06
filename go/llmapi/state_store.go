package llmapi

import (
	"encoding/json"
	"os"
	"sync"
)

// StateStore provides an in-memory and persistent key-value store for agent memories
type StateStore struct {
	mu       sync.RWMutex
	filePath string
	memory   map[string]interface{}
}

// NewStateStore initializes a new StateStore instance
func NewStateStore(filePath string) *StateStore {
	store := &StateStore{
		filePath: filePath,
		memory:   make(map[string]interface{}),
	}
	store.load()
	return store
}

func (s *StateStore) load() {
	if s.filePath == "" {
		return
	}
	data, err := os.ReadFile(s.filePath)
	if err == nil {
		var parsed map[string]interface{}
		if json.Unmarshal(data, &parsed) == nil {
			s.memory = parsed
		}
	}
}

func (s *StateStore) save() {
	if s.filePath == "" {
		return
	}
	data, err := json.MarshalIndent(s.memory, "", "  ")
	if err == nil {
		_ = os.WriteFile(s.filePath, data, 0644)
	}
}

// Get retrieves a value by key
func (s *StateStore) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.memory[key]
	return val, ok
}

// Set stores a key-value pair and persists it
func (s *StateStore) Set(key string, val interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.memory[key] = val
	s.save()
}

// GetAll returns a copy of all stored memories
func (s *StateStore) GetAll() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	copied := make(map[string]interface{})
	for k, v := range s.memory {
		copied[k] = v
	}
	return copied
}
