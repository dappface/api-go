package resolver

import (
	"context"

	"github.com/dappface/api-go/domain/model"
	"github.com/dappface/api-go/interface/api/server/generated"
	"github.com/dappface/api-go/usecase"
)

type Resolver struct {
	postUsecase usecase.PostUseCase
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{
		r,
	}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Feed(ctx context.Context, first *int, last *int, before *string, after *string) (*model.FeedConnection, error) {
	f, err := r.postUsecase.GetConnection(ctx, first, last, before, after)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func New(postUseCase usecase.PostUseCase) generated.Config {
	return generated.Config{
		Resolvers: &Resolver{
			postUseCase,
		},
	}
}
