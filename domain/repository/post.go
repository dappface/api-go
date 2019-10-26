package repository

import (
	"context"

	"github.com/dappface/api-go/domain/model"
)

type PostRepository interface {
	GetAll(ctx context.Context, first *int, last *int, before *string, after *string) (*[]model.Post, error)
	HasPage(ctx context.Context, cursor string, pageType string) (*bool, error)
}
