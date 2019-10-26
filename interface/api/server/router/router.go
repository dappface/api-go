package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(r http.HandlerFunc, q http.HandlerFunc, l http.HandlerFunc) Router {
	return &router{
		r,
		q,
		l,
	}
}

type Router interface {
	CreateHandler() http.Handler
}

type router struct {
	RootHandler     http.HandlerFunc
	QueryHandler    http.HandlerFunc
	LivenessHandler http.HandlerFunc
}

func (r *router) CreateHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/", r.RootHandler)
	mux.HandleFunc("/query", r.QueryHandler)
	mux.HandleFunc("/liveness", r.LivenessHandler)
	return mux
}
