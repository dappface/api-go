package usecase

import (
	"context"
	"log"

	"github.com/dappface/api-go/domain/model"
	"github.com/dappface/api-go/domain/service"
)

func NewPostUseCase(s service.PostService) PostUseCase {
	return &postUseCase{
		s,
	}
}

type PostUseCase interface {
	GetConnection(ctx context.Context, first *int, last *int, before *string, after *string) (*model.FeedConnection, error)
}

type postUseCase struct {
	s service.PostService
}

func (a *postUseCase) GetConnection(ctx context.Context, first *int, last *int, before *string, after *string) (*model.FeedConnection, error) {
	edges, err := a.s.GetEdges(ctx, first, last, before, after)
	if err != nil {
		log.Printf("Failed to get connection: %+v", err)
		return nil, err
	}

	pageInfo, err := a.s.GetPageInfo(ctx, edges)
	if err != nil {
		log.Printf("Failed to get connection: %+v", err)
		return nil, err
	}

	return &model.FeedConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}, nil
}
