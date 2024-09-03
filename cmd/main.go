// cmd/main.go

package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2" // импортируем Fiber
	"github.com/joho/godotenv"    // импортируем godotenv для загрузки переменных окружения из .env
	"github.com/serlenario/my-service/internal/config"
	"github.com/serlenario/my-service/internal/service"
	"log"
)

func main() {
	// Загружаем переменные окружения из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализируем сервис
	srv := service.NewService(cfg.ServiceName)

	// Выводим название сервиса и сообщение в консоль
	fmt.Printf("Service Name: %s\n", cfg.ServiceName)
	fmt.Printf("Message: %s\n", cfg.Message)

	// Используем метод сервиса для вывода имени сервиса
	srv.PrintName()

	// Создаем новое приложение Fiber
	app := fiber.New()

	// Определяем маршрут GET /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(cfg.Message) // отправляем сообщение из конфигурации
	})

	// Запускаем сервер на порту 8080
	log.Fatal(app.Listen(":8080"))
}
