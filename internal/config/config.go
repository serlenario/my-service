package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

// Config структура для хранения конфигурации
type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME"`
	Message     string `envconfig:"MESSAGE"`
}

// LoadConfig загружает конфигурацию из env
func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func init() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	log.Printf("Конфигурация загружена: ServiceName=%s, Message=%s", cfg.ServiceName, cfg.Message)
}
