package service

import (
	"context"
	"log"

	"github.com/dappface/api-go/domain/model"
	"github.com/dappface/api-go/domain/repository"
)

func NewPostService(r repository.PostRepository) PostService {
	return &postService{
		r,
	}
}

type PostService interface {
	GetEdges(ctx context.Context, first *int, last *int, before *string, after *string) ([]*model.PostEdge, error)
	GetPageInfo(ctx context.Context, edges []*model.PostEdge) (*model.PageInfo, error)
}

type postService struct {
	r repository.PostRepository
}

func (s *postService) GetEdges(ctx context.Context, first *int, last *int, before *string, after *string) ([]*model.PostEdge, error) {
	posts, err := s.r.GetAll(ctx, first, last, before, after)
	if err != nil {
		log.Printf("Failed to get all posts: %+v", err)
		return nil, err
	}
	edges := make([]*model.PostEdge, 0)
	for _, post := range *posts {
		edges = append(edges, &model.PostEdge{
			Cursor: post.ID,
			Node:   &post,
		})
	}
	return edges, nil
}

func (s *postService) GetPageInfo(ctx context.Context, edges []*model.PostEdge) (*model.PageInfo, error) {
	hasNextPage, err := s.getHasNextPage(ctx, edges)
	if err != nil {
		log.Printf("Failed to get PageInfo: %+v", err)
		return nil, err
	}
	hasPrevPage, err := s.getHasPrevPage(ctx, edges)
	if err != nil {
		log.Printf("Failed to get PageInfo: %+v", err)
		return nil, err
	}
	return &model.PageInfo{
		HasNextPage:     *hasNextPage,
		HasPreviousPage: *hasPrevPage,
	}, nil
}

func (s *postService) getHasNextPage(ctx context.Context, edges []*model.PostEdge) (*bool, error) {
	if len(edges) == 0 {
		b := false
		return &b, nil
	}

	c := (edges)[len(edges)-1].Cursor
	b, err := s.r.HasPage(ctx, c, "next")
	if err != nil {
		log.Printf("Failed to get HasNextPage: %+v", err)
		return nil, err
	}
	return b, nil
}

func (s *postService) getHasPrevPage(ctx context.Context, edges []*model.PostEdge) (*bool, error) {
	if len(edges) == 0 {
		b := false
		return &b, nil
	}

	c := (edges)[0].Cursor
	b, err := s.r.HasPage(ctx, c, "prev")
	if err != nil {
		log.Printf("Failed to get HasNextPage: %+v", err)
		return nil, err
	}
	return b, nil
}
