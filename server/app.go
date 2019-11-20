package server

import (
	authHttpDelivery "app/auth/delivery/http"
	"app/auth/repository"
	authUseCase "app/auth/usecase"
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

	authUseCase *authUseCase.UserUseCase
}

func NewApplication() *Application {
	database, err := db.New("postgres", "dbname=dev-db user=dev-user password=password")
	if err != nil {
		log.Fatalf("Cant connect to db because of %s", err)
	}

	userRepo := repository.New(database)

	return &Application{
		db:          database,
		authUseCase: authUseCase.New(userRepo),
	}
}

func (a *Application) getHandler() (http.Handler, error) {
	handler := gin.New()
	handler.RedirectTrailingSlash = false

	handler.Use(
		gin.Recovery(),
	)

	apiHandlers := handler.Group("/api/v1")

	authDelivery := authHttpDelivery.New(a.authUseCase)
	authHandler := apiHandlers.Group("/auth")
	{
		authHandler.POST("/sign-in", authDelivery.SignIn)
		authHandler.POST("/sign-up", authDelivery.SignUp)
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
