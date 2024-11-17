package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func responseWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Response with error 5XX", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, playload interface {}) {
	dat, err := json.Marshal(playload)
	if err != nil {
		log.Printf("Failed to marshal json response %v:", playload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}