package app

import (
	"app/internal/adapters/repository"
	"app/pkg/env"
	"log"
)

func Main() {

	env := env.New("")
	dbName, _ := env.String("POSTGRES_DB")
	dbUser, _ := env.String("POSTGRES_USER")
	dbPassword, _ := env.String("POSTGRES_PASSWORD")

	userRepo, err := repository.NewUserRepositoryPostgres(dbName, dbUser, dbPassword, "")
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer userRepo.CloseConnection()

}
