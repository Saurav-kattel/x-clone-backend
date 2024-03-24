package middleware

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/encoder"
	"x-clone.com/backend/src/utils/validator"
)

type ContextKey string

const UserContextKey ContextKey = "user"

func AuthMiddleware(db *sqlx.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("auth_token_x_clone")
			if authToken == "" {
				encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
					Status: http.StatusUnauthorized,
					Res:    models.Message{Message: "Authorization token not provided"},
				})
				return
			}

			// Validate JWT token
			userData, err := validator.ValidateJwt(authToken)
			if err != nil {
				encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
					Status: http.StatusUnauthorized,
					Res:    models.Message{Message: "Invalid JWT token"},
				})
				return
			}

			// Retrieve user from the database using user ID from JWT token
			userInfo, queryErr := user.GetUserByEmail(db, userData.Email)
			if queryErr != nil {
				if queryErr == sql.ErrNoRows {
					encoder.ResponseWriter(w, http.StatusNotFound, models.ErrorResponse{
						Status: http.StatusNotFound,
						Res:    models.Message{Message: "User not found"},
					})
				} else {
					encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
						Status: http.StatusInternalServerError,
						Res:    models.Message{Message: queryErr.Error()},
					})
				}
				return
			}

			ctx := context.WithValue(r.Context(), UserContextKey, userInfo)
			r = r.WithContext(ctx)
			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}
