package appOld

import (
	"app/internal/controllers"
	"app/internal/middleware"
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
	setupLogging()
	defer logFile.Close()

	conf := config.New("")
	service := setupService(conf)
	router := setupRouter(service)
	startServer(conf, router)
}
func setupLogging() {
	var err error
	logFilename := fmt.Sprintf("./logs/%s.txt", time.Now().Format("2006-01-02_15-04-05"))
	logFile, err = os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла лога: %v", err)
	}
	log.SetOutput(logFile)
}
func setupService(conf *config.Config) *services.Services {
	telegramToken, err := conf.String("telegram/token")
	if err != nil {
		log.Fatalf("Ошибка при получении телеграм токена: %v", err)
	}

	// Создаём и настраиваем объекты
	telegram := telegram.New(&telegram.TelegramConfig{
		Token: telegramToken,
	})

	// Перекидываем объекты в сервис
	return &services.Services{
		Telegram: telegram,
		Config:   conf,
	}
}
func setupRouter(services *services.Services) *routes.Routes {
	controller := controllers.New(services)
	middleware := middleware.New(services)
	router := routes.New(controller, middleware)
	router.Setup()
	return router
}

func startServer(conf *config.Config, router *routes.Routes) {
	addr, err := conf.String("server/addr")
	if err != nil {
		log.Fatalf("Ошибка при получении addr: %v", err)
	}
	srv := server.New(addr, router.Mux)

	log.Println("Started")
	fmt.Println("Started")

	srv.Run()
}
