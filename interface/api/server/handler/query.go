package handler

import (
	"net/http"

	"github.com/dappface/api-go/interface/api/server/generated"
	"github.com/dappface/api-go/interface/api/server/resolver"
	"github.com/dappface/api-go/usecase"
	"github.com/99designs/gqlgen/handler"
)

func NewQueryHandler(postUsecase usecase.PostUseCase) http.HandlerFunc {
	return handler.GraphQL(generated.NewExecutableSchema(resolver.New(postUsecase)))
}
