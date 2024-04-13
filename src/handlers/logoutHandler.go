package handlers

import (
	"net/http"

	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/utils/encoder"
)

func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie := &http.Cookie{
			Name:     "auth_token_x_clone",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
		encoder.ResponseWriter(w, 200, models.SuccessResponse{
			Status: 200,
			Res:    "account deleted successfully",
		})

	}
}
