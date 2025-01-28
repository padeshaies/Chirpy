package main

import (
	"fmt"
	"net/http"
)

var baseBody string = `
<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>
`

func (cfg *apiConfig) handleMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	body := fmt.Sprintf(baseBody, cfg.fileServerHits.Load())
	w.Write([]byte(body))
}
