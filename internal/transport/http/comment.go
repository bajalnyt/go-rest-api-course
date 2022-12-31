package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bajalnyt/go-rest-api-course/internal/comment"
	"github.com/gorilla/mux"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	UpdateComment(context.Context, string, comment.Comment) (comment.Comment, error)
	GetComment(context.Context, string) (comment.Comment, error)
	DeleteComment(context.Context, string) error

	GetComments(context.Context) ([]comment.Comment, error)
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		fmt.Println(err)
		return
	}

	cmt, err := h.Service.PostComment(r.Context(), cmt)
	if err != nil {
		fmt.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		fmt.Print(err)
		return
	}
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComments(w http.ResponseWriter, r *http.Request) {
	cmts, err := h.Service.GetComments(r.Context())
	if err != nil {
		fmt.Print(err)
		return
	}
	if err := json.NewEncoder(w).Encode(cmts); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		fmt.Println(err)
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		fmt.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := h.Service.DeleteComment(r.Context(), id)
	if err != nil {
		fmt.Print(err)
		return
	}
}
