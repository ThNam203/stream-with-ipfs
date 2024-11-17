package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sen1or/lets-live/pkg/logger"
	"sen1or/lets-live/user/config"
	"sen1or/lets-live/user/controllers"
	"sen1or/lets-live/user/handlers"
	"sen1or/lets-live/user/middlewares"
	"sen1or/lets-live/user/repositories"

	// TODO: add swagger _ "sen1or/lets-live/auth/docs"
	"time"

	"github.com/jackc/pgx/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
)

type APIServer struct {
	logger    *zap.Logger
	dbConn    *pgx.Conn // For raw sql queries
	serverURL string
	config    config.Config

	errorHandler  *handlers.ErrorHandler
	healthHandler *handlers.HealthHandler
	userHandler   *handlers.UserHandler

	loggingMiddleware middlewares.Middleware
	corsMiddleware    middlewares.Middleware
}

// TODO: make tls usable
func NewAPIServer(dbConn *pgx.Conn, authServerURL string, cfg config.Config) *APIServer {
	var userRepo = repositories.NewUserRepository(dbConn)
	var userCtrl = controllers.NewUserController(userRepo)
	var userHandler = handlers.NewUserHandler(userCtrl)

	var logger, _ = zap.NewProduction()

	return &APIServer{
		logger:    logger,
		dbConn:    dbConn,
		serverURL: authServerURL,

		errorHandler:  handlers.NewErrorHandler(),
		healthHandler: handlers.NewHeathHandler(),
		userHandler:   userHandler,

		loggingMiddleware: middlewares.NewLoggingMiddleware(logger),
		corsMiddleware:    middlewares.NewCORSMiddleware(),
	}
}

func (a *APIServer) ListenAndServe(useTLS bool) {
	server := &http.Server{
		Addr:         a.serverURL,
		Handler:      a.getHandler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if useTLS {
			if _, err := os.Stat(a.config.SSL.ServerCrtFile); err != nil {
				log.Panic("error cant get server cert file", err.Error())
			}

			if _, err := os.Stat(a.config.SSL.ServerKeyFile); err != nil {
				log.Panic("error cant get server key file", err.Error())
			}

			log.Panic("server ends: ", server.ListenAndServeTLS(a.config.SSL.ServerCrtFile, a.config.SSL.ServerKeyFile))
		} else {
			log.Panic("server ends: ", server.ListenAndServe())
		}
	}()

	log.Printf("server running on addr: %s", a.serverURL)
	<-quit

	// Shutdown gracefully
	logger.Infow("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Errorf("server shutdown failed: %+v", err)
	}

	logger.Infow("server exited gracefully")
}

// @title           Let's Live API
// @version         0.1
// @description     The server API

// @contact.name   Nam Huynh
// @contact.email  hthnam203@gmail.com

// @host      localhost:8000
// @BasePath  /v1
func (a *APIServer) getHandler() http.Handler {
	sm := http.NewServeMux()

	sm.HandleFunc("GET /v1/user/{id}", a.userHandler.GetUserByID)
	sm.HandleFunc("POST /v1/user/{id}", a.userHandler.UpdateUser)

	sm.HandleFunc("GET /v1/health", a.healthHandler.GetHealthyState)

	sm.HandleFunc("GET /v1/swagger", httpSwagger.Handler(
		httpSwagger.URL(a.serverURL+"/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	sm.HandleFunc("GET /", a.errorHandler.RouteNotFoundHandler)

	finalHandler := a.corsMiddleware.GetMiddleware(sm)
	finalHandler = a.loggingMiddleware.GetMiddleware(finalHandler)

	return finalHandler
}