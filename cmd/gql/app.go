package main

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/kecci/go-gql-microservice/graph/generated"
	"github.com/kecci/go-gql-microservice/internal/config"
	"github.com/kecci/go-gql-microservice/internal/resolver"

	httpClient "github.com/kecci/go-gql-microservice/internal/outbound/http_client"
	healthrepo "github.com/kecci/go-gql-microservice/internal/repo/health"
	healthsrv "github.com/kecci/go-gql-microservice/internal/service/health"
	"github.com/kecci/go-toolkit/lib/log"
	"github.com/kecci/go-toolkit/lib/sql"
)

func startApp(cfg *config.Config) {
	var (
		ctx = context.Background()
	)

	log.Infoln("Initialize connection to database")
	db, err := sql.Connect(ctx, sql.DBConfig(cfg.Database))
	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Initialize repository Health")
	healthRepo, err := healthrepo.New(db)
	if err != nil {
		log.Fatal(err)
	}

	healthSrv := healthsrv.New(healthRepo)

	httpClient.NewHttpClientOutbound()

	log.Infoln("Initialize Resolver")
	resolver := resolver.NewResolver(
		healthSrv,
	)

	handlerSrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router := newRoutes(handlerSrv)

	startServer(router, cfg)
}
