package app

import (
	adapterUserRepository "app/internal/adapters/user/repository"
	"app/internal/domain/user"
	"app/pkg/database"
	"app/pkg/env"
	"fmt"
	"log"
)

func Main() {

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

	// Закидываем базу в адаптер
	userRepositoryPostgres := adapterUserRepository.NewUserRepositoryPostgres(postgresConnection)

	// Подключаем адаптер в сервис
	userService := user.NewUserService(userRepositoryPostgres)
	_ = userService

	fmt.Println("Hi")
}
