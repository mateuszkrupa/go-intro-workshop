package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	store := &Store{}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	mux.HandleFunc("POST /check", handleCheck(store))
	mux.HandleFunc("GET /history", handleHistory(store))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Link Checker listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}

func handleCheck(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Phase 2 — decode the JSON request body into CheckRequest
		var req CheckRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		if len(req.URLs) == 0 {
			writeError(w, http.StatusBadRequest, "urls must not be empty")
			return
		}

		// Phase 2 — check each URL sequentially and collect results
		// TODO: loop over req.URLs, call checkURL(url) for each,
		//       and append the result to a []CheckResult slice.
		var results []CheckResult
		_ = results

		// Phase 3 — replace the for-loop above with concurrent goroutines.
		// Spawn one goroutine per URL, collect results via a channel,
		// use sync.WaitGroup to wait for all goroutines to finish.
		// Hint: ch := make(chan CheckResult, len(req.URLs))

		// Phase 2 — build response, save to store, write JSON
		// TODO: wrap results in a CheckResponse, save to store, encode as JSON
		_ = store
		writeError(w, http.StatusNotImplemented, "not implemented")
	}
}

func handleHistory(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Phase 4 — retrieve history from store and return as JSON
		// TODO: get history, set Content-Type, encode as JSON
		_ = store
		writeError(w, http.StatusNotImplemented, "not implemented")
	}
}

// writeError writes a JSON error response. Already implemented — use it freely.
func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
