package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("Error fetching comment")
	ErrNotImplemented  = errors.New("Not implemented")
)

// Store - all methods needed to operate by interacting with repository layer
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error

	GetComments(context.Context) ([]Comment, error)
}

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Service - all of the logic
type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) GetComments(ctx context.Context) ([]Comment, error) {
	fmt.Println("retrieving all comments")

	cmt, err := s.Store.GetComments(ctx)
	if err != nil {
		fmt.Println(err)
		return []Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	updated, err := s.Store.UpdateComment(ctx, id, cmt)
	if err != nil {
		return Comment{}, err
	}
	return updated, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	err := s.Store.DeleteComment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	inserted, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return inserted, nil
}
