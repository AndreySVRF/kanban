package main

import (
	"fmt"

	"kanban/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yml")
	if err != nil {
		panic(fmt.Errorf("не удалось загрузить конфигурацию: %w", err))
	}

	fmt.Printf("Запуск сервера на порту: %s\n", cfg.AppPort)
	fmt.Printf("Подключение к базе данных: %s:%s\n", cfg.Database.Host, cfg.Database.Port)
}
