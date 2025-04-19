package app

import (
	"app/internal/routes"
	"app/internal/services"
	"app/pkg/config"
	"app/pkg/server"
	"app/pkg/telegram"
	"fmt"
	"log"
	"os"
	"time"
)

var logFile *os.File

func Main() {
	// Настройка лога
	logFile, err := os.OpenFile(
		fmt.Sprintf("./logs/%s.txt", time.Now().Format("2006-01-02_15-04-05")),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0777,
	)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла лога: %v", err)
	}
	log.SetOutput(logFile)
	defer logFile.Close()

	// Тянем настройку
	conf := config.New("")
	addr, err := conf.String("server/addr")
	if err != nil {
		log.Fatalf("Ошибка при получении addr: %v", err)
	}
	telegramToken, err := conf.String("telegram/token")
	if err != nil {
		log.Fatalf("Ошибка при получении телеграм токена: %v", err)
	}

	// Создаём объекты
	telegram := telegram.New(&telegram.TelegramConfig{
		Token: telegramToken,
	})

	// Перекидываем объекты в сервис
	services := &services.Services{
		Telegram: telegram,
		Config:   conf,
	}

	// Настройка роутера
	router := routes.Setup(services)
	fmt.Println("Started")
	log.Println("Started")

	// Запуск сервера
	srv := server.New(addr, router)
	srv.Run()
}
