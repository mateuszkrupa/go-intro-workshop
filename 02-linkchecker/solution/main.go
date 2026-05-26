package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
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
		var req CheckRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if len(req.URLs) == 0 {
			writeError(w, http.StatusBadRequest, "urls field is required and must not be empty")
			return
		}

		if len(req.URLs) > 20 {
			writeError(w, http.StatusBadRequest, "maximum 20 URLs allowed")
			return
		}

		ch := make(chan CheckResult, len(req.URLs))
		var wg sync.WaitGroup

		for _, url := range req.URLs {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				ch <- checkURL(u)
			}(url)
		}

		// Close channel after all goroutines complete.
		go func() {
			wg.Wait()
			close(ch)
		}()

		var results []CheckResult
		for result := range ch {
			results = append(results, result)
		}

		resp := CheckResponse{Results: results}
		store.Add(resp)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			writeError(w, http.StatusInternalServerError, "internal server error")
		}
	}
}

func handleHistory(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		history := store.GetHistory()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(history); err != nil {
			writeError(w, http.StatusInternalServerError, "internal server error")
		}
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
