package server

import (
	"app/db"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

type Application struct {
	httpServer *http.Server
	db         *sql.DB
}

func NewApplication() *Application {
	database, err := db.New("postgres", "dbname=dev-db user=dev-user password=password")
	if err != nil {
		log.Fatalf("Cant connect to db because of %s", err)
	}

	return &Application{
		db: database,
	}
}

func (a *Application) getHandler() (http.Handler, error) {
	handler := gin.New()
	handler.RedirectTrailingSlash = false

	handler.Use(
		gin.Recovery(),
	)

	apiHandlers := handler.Group("/api/v1")

	authHandler := apiHandlers.Group("/auth")
	{
		authHandler.POST("/sign-in")
		authHandler.POST("/sign-up")
		authHandler.POST("/sign-out")
	}

	return handler, nil
}

func (a *Application) Run() error {
	handler, err := a.getHandler()
	if err != nil {
		return errors.Wrapf(err, "Can't initiate handler.")
	}

	a.httpServer = &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		// service connections
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen: %s", err)
		}
	}()
	return nil

}

func (a *Application) Stop() {
	ctx := context.Background()

	if a.httpServer != nil {
		log.Println("Stopping http server")
		shutdownServerCtx, shutdownServerCancel := context.WithTimeout(ctx, 10*time.Second)
		defer shutdownServerCancel()
		err := a.httpServer.Shutdown(shutdownServerCtx)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if err := a.db.Close(); err != nil {
		log.Printf("Failed to close db instance: %s", err)
	}
}
