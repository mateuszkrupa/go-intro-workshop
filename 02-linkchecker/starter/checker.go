package main

import (
	"net/http"
	"time"
)

// CheckRequest represents the incoming JSON body for POST /check.
type CheckRequest struct {
	URLs []string `json:"urls"`
}

// CheckResult represents a single URL check outcome.
type CheckResult struct {
	URL    string `json:"url"`
	Status int    `json:"status"`
	Ms     int    `json:"ms"`
}

// CheckResponse represents the response for POST /check.
type CheckResponse struct {
	Results []CheckResult `json:"results"`
}

// checkURL performs an HTTP GET and returns the result.
// Already implemented — use it in your handler.
func checkURL(url string) CheckResult {
	client := &http.Client{Timeout: 10 * time.Second}
	start := time.Now()
	resp, err := client.Get(url)
	elapsed := int(time.Since(start).Milliseconds())
	if err != nil {
		return CheckResult{URL: url, Status: 0, Ms: elapsed}
	}
	defer resp.Body.Close()
	return CheckResult{URL: url, Status: resp.StatusCode, Ms: elapsed}
}
