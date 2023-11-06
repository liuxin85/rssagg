package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Respondign with 5xx error: ", msg)
	}
	type errResonse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResonse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, paload interface{}) {
	dat, err := json.Marshal(paload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", paload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
