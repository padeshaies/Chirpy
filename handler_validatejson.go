package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

var profaneWords map[string]struct{} = map[string]struct{}{
	"kerfuffle": struct{}{},
	"sharbert":  struct{}{},
	"fornax":    struct{}{},
}

func handleValidateJSon(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errorVal := errorResponse{Error: "Something went wrong"}
		data, _ := json.Marshal(errorVal)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)

		return
	}

	if len(params.Body) > 140 {
		errorVal := errorResponse{Error: "Chirp is too long"}
		data, _ := json.Marshal(errorVal)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)

		return
	}

	words := strings.Split(params.Body, " ")
	for i, word := range words {
		if _, ok := profaneWords[strings.ToLower(word)]; ok {
			words[i] = "****"
		}
	}
	cleanedBody := strings.Join(words, " ")

	responseBody := returnVals{CleanedBody: cleanedBody}
	data, _ := json.Marshal(responseBody)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
