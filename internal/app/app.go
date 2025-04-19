package app

import (
	"app/internal/routes"
	"app/internal/services"
	"app/pkg/config"
	"app/pkg/telegram"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	telegram := telegram.New(&telegram.TelegramConfig{
		Token: telegramToken,
	})

	services := &services.Services{
		Telegram: telegram,
		Config:   conf,
	}

	// Установки роутов и запуск
	router := routes.Setup(services)
	fmt.Println("Start")

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Server error: ", err)
		}
	}()

	// Обработка сигналов для корректного завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
