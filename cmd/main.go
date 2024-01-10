package main

import (
	"github.com/fdrt29/product-app/pkg"
	"github.com/fdrt29/product-app/pkg/handlers"
	"github.com/fdrt29/product-app/pkg/repositories"
	"github.com/fdrt29/product-app/pkg/services"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializeing congigs: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loadigns env variables: %v", err.Error())
	}

	db, err := repositories.NewPostgresDB(repositories.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %v", err.Error())
	}
	repos := repositories.NewLayer(db)
	services := services.NewLayer(repos)
	handlers := handlers.NewLayer(services)
	srv := new(pkg.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %v", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
