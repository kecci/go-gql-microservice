package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kecci/go-gql-microservice/internal/config"
	"github.com/kecci/go-toolkit/lib/log"
)

func startServer(handler http.Handler, cfg *config.Config) {
	listenApp := listenAppServer(handler, cfg)
	monitorApp := listenMonitorServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Info("received shutdown signal. Trying to shutdown gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	errApp, errMon := listenApp.Shutdown(ctx), monitorApp.Shutdown(ctx)
	if errApp != nil || errMon != nil {
		log.Errorln("Failure while shutting down gracefully, errApp: ", errApp, "; errMon: ", errMon)
	}

	log.Info("Shutdown gracefully completed")

	os.Exit(0)

}

func listenAppServer(handler http.Handler, cfg *config.Config) *http.Server {
	srv := &http.Server{
		Addr:         cfg.Server.GQL.Address,
		WriteTimeout: time.Second * time.Duration(cfg.Server.GQL.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(cfg.Server.GQL.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.Server.GQL.IdleTimeout),
		Handler:      handler,
	}

	go func() {
		log.Println("server running and listen on ", cfg.Server.GQL.Address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(err)
		}
	}()

	return srv
}

func listenMonitorServer() *http.Server {
	srv := &http.Server{
		Addr: ":6060",
	}
	go func() {
		log.Println(srv.ListenAndServe())
	}()

	return srv
}
