package services

import (
	"errors"
	"my-go-app/internal/models"
)

// Временное хранилище пользователей
var users = []models.User{
	{ID: 1, Name: "User1"},
	{ID: 2, Name: "User2"},
}

// Получение всех пользователей
func GetAllUsers() []models.User {
	return users
}

// Получение пользователя по ID
func GetUserByID(id int) (*models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("пользователь не найден")
}

// Создание нового пользователя
func CreateUser(user models.User) {
	users = append(users, user)
}
