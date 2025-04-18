package middleware

import (
	"net/http"
)

// Пример промежуточного ПО для аутентификации
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Здесь можно добавить логику аутентификации
		// Например, проверка заголовков или токенов

		// Если аутентификация успешна, передаем управление следующему обработчику
		next.ServeHTTP(w, r)
	})
}
