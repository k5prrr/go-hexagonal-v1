package app

import (
	"app/internal/adapters/api"
	adapterUserRepository "app/internal/adapters/user/repository"
	"app/internal/domain/user"
	"app/pkg/database"
	"app/pkg/env"
	"app/pkg/server"
	"fmt"
	"log"
	"os"
	"time"
)

var logFile *os.File

func Main() {

	// Настраиваем логи
	setupLogging()
	defer logFile.Close()

	// Читаем настройки
	env := env.New("")
	dbName, err := env.String("POSTGRES_DB")
	if err != nil {
		log.Fatalf("failed env get POSTGRES_DB: %v", err)
	}
	dbUser, err := env.String("POSTGRES_USER")
	if err != nil {
		log.Fatalf("failed env get POSTGRES_USER: %v", err)
	}
	dbPassword, err := env.String("POSTGRES_PASSWORD")
	if err != nil {
		log.Fatalf("failed env get POSTGRES_PASSWORD: %v", err)
	}
	dbPort, err := env.String("POSTGRES1_PORT")
	if err != nil {
		log.Fatalf("failed env get POSTGRES1_PORT: %v", err)
	}
	serverAddress, err := env.String("SERVER_ADDRESS")
	if err != nil {
		log.Fatalf("failed env get SERVER_ADDRESS: %v", err)
	}

	// Подключаем базу
	postgresConnection, err := database.NewPostgresConnection(&database.PostgresConfig{
		Name:     dbName,
		User:     dbUser,
		Password: dbPassword,
		Port:     dbPort,
	})
	if err != nil {
		log.Fatalf("failed NewPostgresConnection: %v", err)
	}
	defer database.ClosePostgresConnection(postgresConnection)

	// Закидываем базу в адаптер пользователей
	userRepositoryPostgres := adapterUserRepository.NewUserRepositoryPostgres(postgresConnection)

	// Подключаем адаптер в сервис
	userService := user.NewUserService(userRepositoryPostgres)

	// Сервис подключаем в роуты
	handlers := api.NewHandlers(userService)
	middleware := api.NewMiddleware(userService)
	err = middleware.Setup()
	if err != nil {
		log.Fatalf("failed middleware.Setup: %v", err)
	}
	router := api.NewRoutes(handlers, middleware)
	router.Setup()

	// Роуты подключаем в сервер
	srv := server.New(serverAddress, router.Mux)
	log.Println("Started")
	fmt.Println("Started")
	srv.Run()
	fmt.Println("Stop")
	log.Println("Stop")
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
