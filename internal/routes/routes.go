package routes

import (
	"app/internal/controllers"
	"app/internal/services"
	"net/http"
	//"app/internal/middleware"
)

func Setup(services *services.Services) http.Handler {
	controllers := controllers.New(services)

	mux := http.NewServeMux()
	mux.HandleFunc("/testSpeed", controllers.TestSpeed)
	//mux.HandleFunc("/testAuth", controllers.testAuth)

	return mux
}
