package main

import (
	"fmt"
	"log"

	"kanban/pkg/config"
	"kanban/pkg/db"
)

func main() {
	// Считываем конфигурацию
	cfg, err := config.LoadConfig("configs/config.yml")
	if err != nil {
		panic(fmt.Errorf("не удалось загрузить конфигурацию: %w", err))
	}

	// Подключаемся к базе данных
	dbpool, err := db.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer dbpool.Close()

	fmt.Println("Успешное подключение к базе данных!")
}
