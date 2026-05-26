package main

import "sync"

// Store holds check history with thread-safe access.
type Store struct {
	mu      sync.Mutex
	history []CheckResponse
}

// Phase 1 — prepend response to history, cap at 10 entries.
// Use s.mu.Lock() / s.mu.Unlock() to protect concurrent access.
func (s *Store) Add(response CheckResponse) {
	// TODO
}

// Phase 1 — return a copy of the history slice.
// Use s.mu.Lock() / s.mu.Unlock() to protect concurrent access.
func (s *Store) GetHistory() []CheckResponse {
	// TODO
	return nil
}
