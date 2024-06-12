package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/tweets"
	"x-clone.com/backend/src/utils/encoder"
)

func DeleteCommentHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "DELETE" {
			encoder.ResponseWriter(w, http.StatusMethodNotAllowed, models.ErrorResponse{
				Status: http.StatusMethodNotAllowed,
				Res: models.Message{
					Message: "invalid http method",
				},
			})
			return
		}
		// getting user  data form r.Context
		_, ok := r.Context().Value(middleware.UserContextKey).(*models.User)
		if !ok {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res:    models.Message{Message: "User information not found"},
			})
			return
		}

		commentId := r.URL.Query().Get("c_id")
		replyId := r.URL.Query().Get("r_id")

		if (commentId == "undefined" || commentId == "") && (replyId == "undefined" || replyId == "") {

			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res:    models.Message{Message: "comment or reply id not found, unable to determine task"},
			})
			return
		}

		validCommentId := commentId != "" && commentId != "undefined"
		validReplyId := replyId != "" && replyId != "undefined"

		if validCommentId {
			if err := tweets.DeleteComment(db, commentId); err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: err.Error()},
				})
				return
			}
		} else if validReplyId {
			if err := tweets.DeleteReply(db, replyId); err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: err.Error()},
				})
				return
			}
		} else {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res:    models.Message{Message: "comment or reply id not found, unable to determine task"},
			})
			return
		}

		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    "deleted successfully",
		})

	}
}
