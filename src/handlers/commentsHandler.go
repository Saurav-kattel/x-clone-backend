package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/notification"
	"x-clone.com/backend/src/tweets"
	"x-clone.com/backend/src/utils/decoder"
	"x-clone.com/backend/src/utils/encoder"
)

func CommentHandlers(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			encoder.ResponseWriter(w, http.StatusMethodNotAllowed, models.ErrorResponse{
				Status: http.StatusMethodNotAllowed,
				Res:    models.Message{Message: "invalid http method"},
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

		data, err := decoder.CommentsPayloadDecoder(r)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res:    models.Message{Message: "req payload: " + err.Error()},
			})
			return
		}

		if data.ParentCommentId != nil {
			handleReply(db, w, data, userData)
		} else {
			handleComment(db, w, data, userData)
		}
	}
}

func handleReply(db *sqlx.DB, w http.ResponseWriter, data *models.Comment, user *models.User) {
	hasParentReply := tweets.CheckForExistingReply(db, *data.ParentCommentId)

	if hasParentReply != nil {
		if hasParentReply == sql.ErrNoRows {
			err := tweets.CreateReplies(db, data.Comment, user.Id, data.TweetId, nil, data.RepliedTO, data.CommentId)
			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: err.Error()},
				})
				return
			}

			notificationMsg := fmt.Sprintf("replied to your comment")
			go notification.CreateNotification(db, &user.Id, &data.RepliedTO, &notificationMsg, &data.TweetId, "reply")
		} else {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res:    models.Message{Message: hasParentReply.Error()},
			})
			return
		}
	} else {
		err := tweets.CreateReplies(db, data.Comment, user.Id, data.TweetId, data.ParentCommentId, data.RepliedTO, data.CommentId)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res:    models.Message{Message: err.Error()},
			})
			return
		}
		notificationMsg := fmt.Sprintf("replied to your comment")
		go notification.CreateNotification(db, &user.Id, &data.RepliedTO, &notificationMsg, &data.TweetId, "reply")
	}

	encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
		Status: http.StatusOK,
		Res:    "reply added successfully",
	})
}

func handleComment(db *sqlx.DB, w http.ResponseWriter, data *models.Comment, user *models.User) {
	err := tweets.CreateComment(db, data.Comment, user.Id, data.TweetId)
	if err != nil {
		encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
			Status: http.StatusInternalServerError,
			Res:    models.Message{Message: err.Error()},
		})
		return
	}
	notificationMsg := fmt.Sprintf("commented on your post")
	go notification.CreateNotification(db, &user.Id, &data.RepliedTO, &notificationMsg, &data.TweetId, "comment")

	encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
		Status: http.StatusOK,
		Res:    "comment added successfully",
	})
}
