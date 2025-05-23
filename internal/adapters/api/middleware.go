package api

import (
	"app/internal/domain/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Middleware struct {
	ApiKeys     map[string]string // API Key -> ServiceName
	UserService user.AnyUserService
}

func NewMiddleware(userService user.AnyUserService) *Middleware {
	return &Middleware{
		ApiKeys:     make(map[string]string),
		UserService: userService,
	}
}

func (m *Middleware) Setup() error {
	data, err := os.ReadFile("configs/apiKeys.json")
	if err != nil {
		return fmt.Errorf("failed to read apiKeys.json: %v", err)
	}

	if err := json.Unmarshal(data, &m.ApiKeys); err != nil {
		return fmt.Errorf("failed to parse apiKeys.json: %v", err)
	}

	return nil
}

// Middleware проверки X-API-Key
func (m *Middleware) CheckApiKey(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-Key")
		if apiKey == "" {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, JSONError("Missing API-Key header"), http.StatusUnauthorized)
			return
		}

		serviceName, exists := m.ApiKeys[apiKey]
		if !exists {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, JSONError("Invalid or expired API Key"), http.StatusUnauthorized)
			return
		}

		log.Println("%s %s", serviceName, r.Method)
		next(w, r)
	}
}
