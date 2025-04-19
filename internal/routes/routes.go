package routes

import (
	"net/http"
	"app/internal/controllers"
	"app/internal/services"
	//"app/internal/middleware"
)

func Setup(services *services.Services) http.Handler {
	controllers := controllers.New(services)

	mux := http.NewServeMux()
	mux.HandleFunc("/testSpeed", controllers.TestSpeed)
	//mux.HandleFunc("/testAuth", controllers.testAuth)

	return mux
}
