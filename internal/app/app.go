package app

import (
	"fmt"
	"app/pkg/config"
	"log"
	"os"
	"time"
	"net/http"
	"app/internal/routes"
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
	port, err := config.String("server/port")
	if err != nil {
		log.Fatalf("Ошибка при получении порта: %v", err)
	}

	telegram := Telegram.New(&Telegram.TelegramConfig{
		Token: "123",
	})

	// Установки роутов и запуск
	router := routes.SetupRoutes()
	fmt.Println("Start")
	log.Fatalf("Ошибка серверв %v", http.ListenAndServe(port, router))
}
