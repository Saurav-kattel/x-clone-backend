package handlers

import (
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/tweets"
	"x-clone.com/backend/src/utils/decoder"
	"x-clone.com/backend/src/utils/encoder"
)

func CommentHandlers(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			encoder.ResponseWriter(w, http.StatusMethodNotAllowed, models.ErrorResponse{
				Status: http.StatusMethodNotAllowed,
				Res: models.Message{
					Message: "invalid http method",
				},
			})
			return
		}

		userData, ok := r.Context().Value(middleware.UserContextKey).(*models.User)
		if !ok {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res:    models.Message{Message: "User information not found"},
			})
			return
		}

		// parsing req.FormValue
		data, err := decoder.CommentsPayloadDecoder(r)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "req payload: " + err.Error(),
				},
			})
			return
		}

		if data.ParentCommentId != nil {
			hasParentReply := tweets.CheckForExistingReply(db, *data.ParentCommentId)
			// checking if the existing reply for error
			if hasParentReply != nil {
				// checking for not found error
				if hasParentReply == sql.ErrNoRows {
					//if no rows create a reply with parent comment id nil
					err := tweets.CreateReplies(db, data.Comment, userData.Id, data.TweetId, nil, data.RepliedTO, data.CommentId)
					if err != nil {
						encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
							Status: http.StatusInternalServerError,
							Res: models.Message{
								Message: err.Error(),
							},
						})
						return
					}
				} else {
					encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
						Status: http.StatusInternalServerError,
						Res: models.Message{
							Message: hasParentReply.Error(),
						},
					})
					return
				}
			} else {
				err := tweets.CreateReplies(db, data.Comment, userData.Id, data.TweetId, data.ParentCommentId, data.RepliedTO, data.CommentId)
				if err != nil {
					encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
						Status: http.StatusInternalServerError,
						Res: models.Message{
							Message: err.Error(),
						},
					})
					return
				}
			}
		}

		if data.ParentCommentId == nil {
			err := tweets.CreateComment(db, data.Comment, userData.Id, data.TweetId)
			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: err.Error(),
					},
				})
				return
			}
		}

		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    "comments added successfully",
		})

	}
}
