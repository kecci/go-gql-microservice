package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func newRoutes(handlerSrv *handler.Server) *chi.Mux {
	router := chi.NewRouter()

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", handlerSrv)

	return router
}
