package main

import (
	"foveo/internal/config"
	"foveo/internal/handler"
	"foveo/internal/repository"
	"foveo/internal/router"
	"foveo/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	db, err := sqlx.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// init repo
	userRepo := repository.InitUserRepo(db)

	// init service
	userService := service.InitUserService(userRepo)

	// init handler
	userHandler := handler.InitUserHandler(userService)
	handler := handler.NewHandler(userHandler)

	log.Println("Server started on :8090")
	http.ListenAndServe(":8090", router.NewRouter(handler))

}
