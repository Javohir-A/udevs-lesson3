package main

import (
	app "github.com/udevs/lesson3/api"
	"github.com/udevs/lesson3/api/handlers"
	"github.com/udevs/lesson3/config"
	"github.com/udevs/lesson3/mongo"
	"github.com/udevs/lesson3/pkg/logger"
	"github.com/udevs/lesson3/storage"
	"go.uber.org/zap"
)

func main() {
	logger.Initialize()
	log := logger.GetLogger()

	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed to get config", zap.Error(err))
	}
	mongoDB, err := mongo.Connect(&cfg.MongoDB)
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}
	testDB := mongoDB.Database("test_db")

	productsCollection := testDB.Collection("products")
	ordersCollection := testDB.Collection("orders")

	productStorage := storage.NewProductStorage(productsCollection)
	orderStorage := storage.NewOrdersStorage(ordersCollection)

	proHandler := handlers.NewProductsHandler(productStorage, log)
	ordHandler := handlers.NewOrdersHandler(orderStorage, log)

	httpservice := app.NewHttpService(ordHandler, proHandler, log, cfg)

	httpservice.Run()
}
