package main

import (
	"encoding/json"
	"log"
	"net/http"
)
type successful struct{
	Msg string `json:"msg"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal the JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	// Logging on server-side
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg)
	}
	// Reflecting the error to the client
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{Error: msg})
}

func handleRoot(w http.ResponseWriter, r *http.Request){
	respondWithJSON(w, 200, successful{Msg: "Welcome, Please provide only .xml links as feeds."})
}