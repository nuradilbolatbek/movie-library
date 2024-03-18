package main

import (
	"VkTask"
	"VkTask/package/handler"
	"VkTask/package/repository"
	"VkTask/package/service"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	srv := new(VkTask.Server)

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	//handlers := handler.NewHandler(services)

	authService := service.NewAuthService(repos.Authorization)
	handlers := handler.NewHandler(services, authService)

	if err := srv.Run("8080", handlers); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
