package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/serlenario/my-service/internal/config"
	"github.com/serlenario/my-service/internal/service"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
	// Загружаем конфигурацию
	cfg, err := internal.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализируем сервис
	srv := service.NewService(cfg.ServiceName)

	// Выводим название сервиса в консоль
	fmt.Printf("Service Name: %s\n", cfg.ServiceName)
	fmt.Printf("Message: %s\n", cfg.Message)

	// Используем метод сервиса
	srv.PrintName()
}
