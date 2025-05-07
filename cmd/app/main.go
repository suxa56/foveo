package main

import (
	"foveo/internal/config"
	"foveo/internal/handler"
	"foveo/internal/repository"
	"foveo/internal/router"
	"foveo/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	db, err := sqlx.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		log.Error("db error", slog.String("error", err.Error()))
	}
	defer db.Close()

	// init repo
	userRepo := repository.InitUserRepo(db)

	// init service
	userService := service.InitUserService(userRepo)

	// init handler
	userHandler := handler.InitUserHandler(userService)
	handler := handler.NewHandler(userHandler)

	log.Info("Server started on :8090")
	http.ListenAndServe(":8090", router.NewRouter(handler))

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	}
	return log
}
