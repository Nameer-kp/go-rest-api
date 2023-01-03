package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchComment   = errors.New("failed to fetch Comment by id")
	ErrNotImplemented = errors.New("not implemented")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(ctx context.Context, cmt Comment) (Comment, error)
	DeleteComment(ctx context.Context, id string) error
	UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error)
}
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
		return Comment{}, ErrFetchComment
	}
	return cmt, nil
}
func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	updatedCmt, err := s.Store.UpdateComment(ctx, id, cmt)
	if err != nil {
		return Comment{}, err
	}
	return updatedCmt, nil
}
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}
