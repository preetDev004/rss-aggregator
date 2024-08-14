package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/preetDev004/rss-aggregator/db"
)
const MIN_CHARS = 2
const MAX_CHARS = 15

type parameters struct{
	Name string `json:"name"`
}

func decodeParams(r *http.Request) (parameters, error){
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	return params, err
}

// creating handler for apiConfig struct so we can use queries functions inside the handler. 
func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request){

	params, err := decodeParams(r)
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %v", err))
		return
	}
	if len(params.Name) <= 2 || len(params.Name) > 15 {
		respondWithError(w, 400, fmt.Sprintf("Please enter a name with more than %v characters and less than %v characters!", MIN_CHARS, MAX_CHARS))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID : uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err!=nil{
		respondWithError(w, 500, fmt.Sprintf("Error Creating the user: %v", err))
		return
	}
	respondWithJSON(w, 201, dbUserToUser(user))
}