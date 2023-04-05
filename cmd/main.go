package main

import (
	"log"
	app "webapp"
	"webapp/pkg/handler"
	"webapp/pkg/repository"
	"webapp/pkg/service"

	_ "github.com/lib/pq"
)

func main() {
	db, err := repository.NewPostgresDb(&repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "tez_taxi",
		DBname:   "tez_taxi",
		SSLMmode: "disable",
		Password: "secret",
	})

	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(app.Server)

	err = srv.Run("8080", handler.InitRoutes())
	if err != nil {
		log.Fatal()
	}
	
	return
}
