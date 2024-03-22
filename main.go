package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"x-clone.com/backend/src/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")
	db, dbErr := database.ConnectDB(username, password, dbName)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer db.Close()

	log.Print("server connected with db")

	if err := http.ListenAndServe("127.0.0.1:"+port, nil); err != nil {
		log.Fatal(err)
	}

}
