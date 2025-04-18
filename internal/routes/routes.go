package routes

import (
	"net/http"
	"app/internal/controllers"
	//"app/internal/middleware"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/testSpeed", controllers.TestSpeed)
	//mux.HandleFunc("/testAuth", controllers.testAuth)

	return mux
}
