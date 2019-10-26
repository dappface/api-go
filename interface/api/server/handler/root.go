package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
)

func NewRootHandler() http.HandlerFunc {
	return handler.Playground("GraphQL playground", "/query")
}
