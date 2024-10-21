package main

import (
	orderv23 "gofermart/internal/api/handlers/orderv2"
	"gofermart/internal/repository/kafka/orderv2"
	"gofermart/internal/repository/worker/orders"
	orderv22 "gofermart/internal/service/orderv2"
	"gofermart/pkg/kafka"
	"log"
	"net/http"

	"gofermart/internal/api"
	"gofermart/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	//db, err := mongo.NewMongoDB(cfg.DatabaseURI, cfg.DatabaseName)
	//if err != nil {
	//	log.Fatalf("Ошибка подключения к базе данных: %v", err)
	//}
	//defer db.Close()
	//
	//// Инициализация базы данных
	//if err := database.InitDatabase(db.Database()); err != nil {
	//	log.Fatalf("Ошибка инициализации базы данных: %v", err)
	//}

	//userRepository := userRepo.NewRepository(db.Database())
	//orderRepository := orderRepo.NewRepository(db.Database())
	//balanceRepository := balanceRepo.NewRepository(db.Database())
	//
	//userSvc := userService.NewService(userRepository)
	//orderSvc := orderService.NewService(orderRepository)
	//balanceSvc := balanceService.NewService(balanceRepository)

	kafkaProducer, err := kafka.NewProducerClient([]string{cfg.BrokersList})
	if err != nil {
		panic(err)
	}

	repository := orderv2.NewRepository(kafkaProducer)
	service := orderv22.NewService(repository)
	handler := orderv23.NewHandler(service)

	router := api.NewRouter(handler)

	kafkaConsumer, err := kafka.NewKafkaConsumer([]string{cfg.BrokersList}, "0", []string{"orders"}, &kafka.ExampleHandler{})
	if err != nil {
		panic(err)
	}

	worker := orders.NewWorker(kafkaConsumer)
	go worker.Start()

	log.Printf("Запуск сервера на %s", cfg.RunAddress)
	if err := http.ListenAndServe(cfg.RunAddress, router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
