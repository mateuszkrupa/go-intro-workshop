package main

import "sync"

// Store holds check history with thread-safe access.
type Store struct {
	mu      sync.Mutex
	history []CheckResponse
}

// Add prepends a CheckResponse to the history and caps it at 10 entries.
func (s *Store) Add(response CheckResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.history = append([]CheckResponse{response}, s.history...)
	if len(s.history) > 10 {
		s.history = s.history[:10]
	}
}

// GetHistory returns a copy of the history slice.
func (s *Store) GetHistory() []CheckResponse {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]CheckResponse, len(s.history))
	copy(result, s.history)
	return result
}
