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

	authStack := middleware.CreateStack(
		middleware.Logger,
		middleware.AuthMiddleware(db),
	)

	unAuthStack := middleware.CreateStack(
		middleware.Logger,
	)

	router.Handle("/api/v1/user/register", unAuthStack(http.HandlerFunc(handlers.RegisterUserHandler(db))))
	router.Handle("/api/v1/user/login", unAuthStack(http.HandlerFunc(handlers.LoginUserHandler(db))))
	router.Handle("/api/v1/user/verify/email", unAuthStack(http.HandlerFunc(handlers.VerifyEmailHandler(db))))
	router.Handle("/api/v1/user/account/forgot-password", unAuthStack(http.HandlerFunc(handlers.UpdateForgottenPasswordHandler(db))))

	router.Handle("/api/v1/user/account/image", authStack(http.HandlerFunc(handlers.InsertProfileHandler(db))))
	router.Handle("/api/v1/user/account/delete", authStack(http.HandlerFunc(handlers.DeleteUserAccountHandler(db))))
	router.Handle("/api/v1/user/account/username", authStack(http.HandlerFunc(handlers.UpdateUsernameHandler(db))))
	router.Handle("/api/v1/user/account/password", authStack(http.HandlerFunc(handlers.UpdatePasswordHandler(db))))
	router.Handle("/api/v1/tweet/post", authStack(http.HandlerFunc(handlers.CreateTweetHandler(db))))
	router.Handle("/api/v1/tweet/delete", authStack(http.HandlerFunc(handlers.DeleteTweetHandler(db))))
	server := http.Server{
		Addr:    "localhost:4000",
		Handler: router,
	}
	//listening server on localhost
	if err := server.ListenAndServe(); err != nil {
		log.Panic("Error starting HTTP server:", err)
	}
}
