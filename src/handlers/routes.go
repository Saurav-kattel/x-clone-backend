package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
)

func Routers(db *sqlx.DB) *http.ServeMux {
	router := http.NewServeMux()

	authStack := middleware.CreateStack(
		middleware.Logger,
		middleware.AuthMiddleware(db),
	)

	unAuthStack := middleware.CreateStack(
		middleware.Logger,
	)

	router.Handle("/api/v1/user/register", unAuthStack(RegisterUserHandler(db)))
	router.Handle("/api/v1/user/login", unAuthStack(LoginUserHandler(db)))
	router.Handle("/api/v1/user/verify/email", unAuthStack(VerifyEmailHandler(db)))
	router.Handle("/api/v1/user/account/forgot-password", unAuthStack(UpdateForgottenPasswordHandler(db)))
	router.Handle("/api/v1/user/account/logout", unAuthStack(LogoutHandler()))

	router.Handle("/api/v1/tweet/author/image", unAuthStack(GetAuthorImageHandler(db)))
	router.Handle("/api/v1/tweet/get", unAuthStack(GetTweetsHandler(db)))
	router.Handle("/api/v1/tweet/image", unAuthStack(GetTweetImageHandler(db)))
	router.Handle("/api/v1/tweet/like/count", unAuthStack(LikesCountHandler(db)))

	router.Handle("/api/v1/user/account/image", authStack(InsertProfileHandler(db)))
	router.Handle("/api/v1/user/account/delete", authStack(DeleteUserAccountHandler(db)))
	router.Handle("/api/v1/user/account/get-image", authStack(GetProfileImageHandler(db)))
	router.Handle("/api/v1/user/account/username", authStack(UpdateUsernameHandler(db)))
	router.Handle("/api/v1/user/account/password", authStack(UpdatePasswordHandler(db)))
	router.Handle("/api/v1/user/get", authStack(GetUserHandler(db)))

	router.Handle("/api/v1/tweet/post", authStack(CreateTweetHandler(db)))
	router.Handle("/api/v1/tweet/delete", authStack(DeleteTweetHandler(db)))
	router.Handle("/api/v1/tweet/like", authStack(TweetLikeHandler(db)))
	return router
}
