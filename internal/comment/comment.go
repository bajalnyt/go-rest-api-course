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

func (s *Service) UpdateComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, id string) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
