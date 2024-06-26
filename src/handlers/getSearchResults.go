package handlers

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/encoder"
)

func GetUsersSearchResult(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			encoder.ResponseWriter(w, http.StatusMethodNotAllowed, models.ErrorResponse{
				Status: http.StatusMethodNotAllowed,
				Res: models.Message{
					Message: "invalid http method",
				},
			})
			return
		}

		query := r.URL.Query().Get("query")

		if query == "undefined" || query == "null" {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "query is undefined",
				},
			})
			return
		}

		query = fmt.Sprintf("%%%s%%", query)
		data, err := user.SearchResult(db, query)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res: models.Message{
					Message: err.Error(),
				},
			})
			return
		}

		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    data,
		})

	}
}
