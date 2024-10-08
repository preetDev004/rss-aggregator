package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/preetDev004/rss-aggregator/db"
)

func main() {
	// Load the environment variables
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No PORT found!")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Database URL not found!")
	}

	// connecting to database
	connection := connectToDB(dbURL)
	defer closeDB(connection) // defer closing the database

	db := db.New(connection) // Create queries from db to query the database
	apiCfg := apiConfig{          // creating struct.
		DB: db,
	}
	// start scraping the feeds in a saparate go-routine
	go startScraping(db, 10, time.Minute)

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

	v1Router.Get("/", handleRoot)
	// user
	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))

	// feed
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	v1Router.Get("/feeds", apiCfg.handleGetAllFeeds)

	// feedFollows
	v1Router.Post("/feedFollow", apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollow))
	v1Router.Get("/feedFollow", apiCfg.middlewareAuth(apiCfg.handleGetUserFeedFollows))
	v1Router.Delete("/feedFollow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handleDeleteUserFeedFollow))

	// posts
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handleGetPosts))

	router.Mount("/v1", v1Router) // Mount the nested router to the main one

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
