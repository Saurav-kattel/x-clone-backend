package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"x-clone.com/backend/src/database"
	"x-clone.com/backend/src/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := database.ConnectDB(username, password, dbName)
	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}
	defer db.Close()

	log.Print("Server connected with the database.")

	registerUser := handlers.RegisterUserHandler(db)

	http.Handle("/api/v1/user/register", registerUser)

	if err := http.ListenAndServe("127.0.0.1:"+port, nil); err != nil {
		log.Panic("Error starting HTTP server:", err)
	}
}
