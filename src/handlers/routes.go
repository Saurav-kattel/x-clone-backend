package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
)

func Routers(db *sqlx.DB) *http.ServeMux {
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

	router.Handle("/api/v1/user/register", unAuthStack(http.HandlerFunc(RegisterUserHandler(db))))
	router.Handle("/api/v1/user/login", unAuthStack(http.HandlerFunc(LoginUserHandler(db))))
	router.Handle("/api/v1/user/verify/email", unAuthStack(http.HandlerFunc(VerifyEmailHandler(db))))
	router.Handle("/api/v1/user/account/forgot-password", unAuthStack(http.HandlerFunc(UpdateForgottenPasswordHandler(db))))
	router.Handle("/api/v1/user/account/logout", unAuthStack(http.HandlerFunc(LogoutHandler())))

	router.Handle("/api/v1/user/account/image", authStack(http.HandlerFunc(InsertProfileHandler(db))))
	router.Handle("/api/v1/user/account/delete", authStack(http.HandlerFunc(DeleteUserAccountHandler(db))))
	router.Handle("/api/v1/user/account/get-image", authStack(http.HandlerFunc(GetProfileImageHandler(db))))
	router.Handle("/api/v1/user/account/username", authStack(http.HandlerFunc(UpdateUsernameHandler(db))))
	router.Handle("/api/v1/user/account/password", authStack(http.HandlerFunc(UpdatePasswordHandler(db))))
	router.Handle("/api/v1/user/get", authStack(http.HandlerFunc(GetUserHandler(db))))

	router.Handle("/api/v1/tweet/post", authStack(http.HandlerFunc(CreateTweetHandler(db))))
	router.Handle("/api/v1/tweet/delete", authStack(http.HandlerFunc(DeleteTweetHandler(db))))
	return router

}
