package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lovelymod/task-management-backend/internal/bootstrap"
	"github.com/lovelymod/task-management-backend/internal/handler"
	"github.com/lovelymod/task-management-backend/internal/repository"
	"github.com/lovelymod/task-management-backend/internal/router"
	"github.com/lovelymod/task-management-backend/internal/usecase"
)

func main() {
	r := gin.Default()

	app := bootstrap.AppInit()
	defer func() {
		if err := app.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	authRepo := repository.NewAuthHandler(app.Mc)
	authUsecase := usecase.NewAuthUsecase(authRepo, time.Second*5, app.Config)
	authHandler := handler.NewAuthHandler(authUsecase)

	handlers := router.Handlers{
		AuthHandler: authHandler,
	}

	router.SetupRouter(r, &handlers)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
