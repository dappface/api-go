package registry

import (
	"net/http"

	"github.com/dappface/api-go/domain/service"
	"github.com/dappface/api-go/infrastructure/firestore"
	"github.com/dappface/api-go/interface/api/server/handler"
	"github.com/dappface/api-go/interface/api/server/router"
	"github.com/dappface/api-go/usecase"
)

type Registry interface {
	NewHandler() http.Handler
}

type registry struct{}

func NewRegistry() Registry {
	return &registry{}
}

func (s *registry) NewHandler() http.Handler {
	// repository
	postRepository := firestore.NewPostRepository()

	// service
	postService := service.NewPostService(postRepository)

	// usecase
	postUseCase := usecase.NewPostUseCase(postService)

	// handler
	rootHandler := handler.NewRootHandler()
	queryHandler := handler.NewQueryHandler(postUseCase)
	livenessHandler := handler.NewLivenessHandler()

	r := router.NewRouter(rootHandler, queryHandler, livenessHandler)
	return r.CreateHandler()
}
