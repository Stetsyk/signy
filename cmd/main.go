package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Stetsyk/signy"
	"github.com/Stetsyk/signy/pkg/handler"
	"github.com/Stetsyk/signy/pkg/repository"
	"github.com/Stetsyk/signy/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	db, err := repository.NewMysqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	createTables(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// handlers := new(handler.Handler)
	srv := new(signy.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Signy App Started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
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
