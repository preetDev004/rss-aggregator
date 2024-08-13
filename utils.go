package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/preetDev004/rss-aggregator/db"
)

type apiConfig struct{
	DB *db.Queries
}

func connectToDB(dbURL string) *sql.DB{
	connection, err := sql.Open("postgres", dbURL)
    if err != nil {
        panic(err)
    }
	fmt.Println("connected to database")
	return connection
}

func closeDB(db *sql.DB){
	log.Println("Closing Connection to the Database.")
	db.Close()
}