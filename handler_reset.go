package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	cfg.fileServerHits.Store(0)

	err := cfg.dbQueries.Reset(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error deleting users"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Environment reset"))
}
