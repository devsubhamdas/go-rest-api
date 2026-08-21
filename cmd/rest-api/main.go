package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Subham-Das-98/go-rest-api/internal/config"
	"github.com/Subham-Das-98/go-rest-api/internal/handler"
	"github.com/Subham-Das-98/go-rest-api/internal/models"
	"github.com/Subham-Das-98/go-rest-api/internal/repository"
	"github.com/Subham-Das-98/go-rest-api/internal/service"
	"github.com/Subham-Das-98/go-rest-api/internal/storage"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database connection
	db, err := storage.NewPostgres(cfg.Postgres)
	if err != nil {
		slog.Error("failed to connect database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// migration
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		slog.Error("database migration failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	// Repository
	userRepository := repository.NewUserRepository(db)

	// Service
	userService := service.NewUserService(userRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService)

	// setup route
	router := http.NewServeMux()

	router.HandleFunc("GET /api/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome go-rest-api"))
	})

	router.HandleFunc(
		"POST /api/users",
		userHandler.CreateUser,
	)

	router.HandleFunc(
		"GET /api/users",
		userHandler.GetUsers,
	)

	router.HandleFunc(
		"GET /api/users/{id}",
		userHandler.GetUser,
	)

	router.HandleFunc(
		"PUT /api/users/{id}",
		userHandler.UpdateUser,
	)

	router.HandleFunc(
		"DELETE /api/users/{id}",
		userHandler.DeleteUser,
	)

	// start server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")
}
