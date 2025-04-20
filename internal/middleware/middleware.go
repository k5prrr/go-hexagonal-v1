package middleware

import (
	"app/internal/services"
	"log"
	"net/http"
)

type Middleware struct {
	Services *services.Services
}

func New(services *services.Services) *Middleware {
	return &Middleware{
		Services: services,
	}
}

// Auth теперь возвращает http.HandlerFunc вместо http.Handler
func (m *Middleware) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("userID")
		userKEY := r.Header.Get("userKEY")

		if userID == "" || userKEY == "" {
			log.Printf("Unauthorized request - missing headers")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !m.isValidUser(userID, userKEY) {
			log.Printf("Unauthorized request - invalid credentials")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}

func (m *Middleware) isValidUser(userID, userKEY string) bool {
	// В реальном приложении здесь должна быть проверка через Services
	return userKEY == "valid_key_"+userID
}
