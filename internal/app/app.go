package app

import (
	"log"
	"net/http"
)

func NewApp(httpServer *HttpServer) App {
	return &app{httpServer: httpServer}
}

type App interface {
	Start() error
}

type app struct {
	httpServer *HttpServer
}

func (a *app) Start() error {
	if err := http.ListenAndServe(":8080", a.httpServer.Handler); err != nil {
		log.Fatalf("Failed to start server %v", err)
		// os.Exit(1)
		return err
	}
	return nil
}
