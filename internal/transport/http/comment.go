package http

import (
	"context"
	"encoding/json"
	"github.com/Nameer-kp/go-rest-api/internal/comment"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(context.Context, string) (comment.Comment, error)
	UpdateComment(context.Context, string, comment.Comment) (comment.Comment, error)
	DeleteComment(context.Context, string) error
}

type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func convertPostCommentRequestToComment(c PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   c.Slug,
		Author: c.Author,
		Body:   c.Body,
	}
}
func convertCommentToPostCommentRequest(c comment.Comment) PostCommentRequest {
	return PostCommentRequest{
		Slug:   c.Slug,
		Author: c.Author,
		Body:   c.Body,
	}
}
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}
	validate := validator.New()
	if err := validate.Struct(cmt); err != nil {
		http.Error(w, "not a valid comment", http.StatusBadRequest)
		return
	}
	convertedCmt := convertPostCommentRequestToComment(cmt)
	postedComment, err := h.Service.PostComment(r.Context(), convertedCmt)
	if err != nil {
		log.Println("Error while Posting comment")
		return
	}
	if err := json.NewEncoder(w).Encode(postedComment); err != nil {
		log.Println("error while sending response (posting comment)")
		return
	}
}
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(cmt)
	if err != nil {
		log.Println("error while sending response(reading comment)")
		return
	}

}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.Service.DeleteComment(r.Context(), id)
	if err != nil {
		log.Println("error deleting comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Println("Error while updating comment")
		return
	}
	err = json.NewEncoder(w).Encode(cmt)
	if err != nil {
		log.Println("error while sending response(updating comment)")
		return
	}
}
