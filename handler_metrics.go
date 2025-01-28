package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handleMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	body := fmt.Sprintf("Hits: %d\n", cfg.fileServerHits.Load())
	w.Write([]byte(body))
}
