package main

import (
	"github.com/Stetsyk/signy"
	"github.com/Stetsyk/signy/pkg/handler"
	"github.com/Stetsyk/signy/pkg/repository"
	"github.com/Stetsyk/signy/pkg/service"
	"github.com/sirupsen/logrus"
)

// "fmt"
// "github.com/gin-gonic/gin"

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// handlers := new(handler.Handler)
	srv := new(signy.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}
