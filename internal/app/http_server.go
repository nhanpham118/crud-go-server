package app

import (
	"crud-go-server/internal/pkg/api"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHttpServer(
	studentHandler *api.StudentHandler,
	moduleHandler *api.ModuleHandler,
	markHandler *api.MarkHandler,
) *HttpServer {
	router := chi.NewRouter()

	// Use logger on request
	router.Use(middleware.Logger)

	// Mount path requests
	router.Mount("/student", studentHandler.Route())
	router.Mount("/module", moduleHandler.Route())
	router.Mount("/mark", markHandler.Route())

	server := &http.Server{
		Handler: router,
	}
	return &HttpServer{server}
}

type HttpServer struct {
	*http.Server
}
