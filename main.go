package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"x-clone.com/backend/src/database"
	"x-clone.com/backend/src/handlers"
	"x-clone.com/backend/src/middleware"
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

	// router
	router := http.NewServeMux()

	// Register handlers for each endpoint separately
	router.HandleFunc("/api/v1/user/register", handlers.RegisterUserHandler(db))
	router.HandleFunc("/api/v1/user/login", handlers.LoginUserHandler(db))
	router.HandleFunc("/api/v1/user/account/image", handlers.InsertProfileHandler(db))
	router.HandleFunc("/api/v1/user/account/delete", handlers.DeleteUserAccountHandler(db))
	router.HandleFunc("/api/v1/user/account/username", handlers.UpdateUsernameHandler(db))
	router.HandleFunc("/api/v1/user/account/password", handlers.UpdatePasswordHandler(db))

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: middleware.Logger(router),
	}
	//listening server on localhost
	if err := server.ListenAndServe(); err != nil {
		log.Panic("Error starting HTTP server:", err)
	}
}
