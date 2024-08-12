package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load the environment variables
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No PORT found!")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// standered practice - nested routing in case you release other versions
	v1Router := chi.NewRouter()
	// v1Router.HandleFunc("/healthz", handleReadiness) - Open for all requests
	v1Router.Get("/healthz", handleReadiness) // open for Get requests only
	v1Router.Get("/error", handleError) // open for Get requests only
	router.Mount("/v1",v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Starting server on port: %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error running the server: ", err)
	}

	
	
}

