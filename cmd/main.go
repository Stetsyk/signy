package main

import (
	"log"

	"github.com/Stetsyk/signy"
	"github.com/Stetsyk/signy/pkg/handler"
	"github.com/Stetsyk/signy/pkg/repository"
	"github.com/Stetsyk/signy/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	db, err := repository.NewMysqlDB(repository.Config{
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "1111",
		DBName:   "documents",
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	createTables(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// handlers := new(handler.Handler)
	srv := new(signy.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func createTables(db *gorm.DB) {
	if db.Migrator().HasTable(&signy.User{}) {
		db.Exec("DROP TABLE users")
	}
	db.AutoMigrate(&signy.User{})
	if db.Migrator().HasTable(&signy.Document{}) {
		db.Exec("DROP TABLE documents")
	}
	db.AutoMigrate(&signy.Document{})
	if (db.Migrator().HasTable(&signy.Signature{})) {
		db.Exec("DROP TABLE signatures")
	}
	db.AutoMigrate(&signy.Signature{})
}
