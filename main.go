package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"x-clone.com/backend/src/database"
	"x-clone.com/backend/src/handlers"
)

func main() {
	//loading environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	//port := os.Getenv("PORT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//connecting to database
	db, err := database.ConnectDB(username, password, dbName)
	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}
	defer db.Close()

	log.Print("Server connected with the database.")

	//http endpoints/routes
	router := handlers.Routers(db)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		Debug:            true,
	})

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: c.Handler(router),
	}
	//listening server on localhost
	if err := server.ListenAndServe(); err != nil {
		log.Panic("Error starting HTTP server:", err)
	}
}
