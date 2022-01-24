package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	auth_server "github.com/th2empty/auth-server"
	"github.com/th2empty/auth-server/pkg/handler"
	"github.com/th2empty/auth-server/pkg/repository"
	"github.com/th2empty/auth-server/pkg/service"
	"os"
)

func main() {
	log.SetFormatter(new(log.JSONFormatter))

	if err := initConfig(); err != nil {
		log.WithFields(log.Fields{
			"package":  "main",
			"file":     "main.go",
			"function": "initConfig",
			"error":    err,
		}).Fatalf("error initializing configs")
	}

	if err := godotenv.Load(); err != nil {
		log.WithFields(log.Fields{
			"package":  "main",
			"file":     "main.go",
			"function": "initConfig",
			"error":    err,
		}).Fatalf("error loading env variables")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"package":  "main",
			"file":     "main.go",
			"function": "main",
			"error":    err,
		}).Fatalf("failed to initialize database")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(auth_server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.WithFields(log.Fields{
			"package":  "main",
			"file":     "main.go",
			"function": "main",
			"error":    err,
		}).Fatalf("error occured while running http server")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
