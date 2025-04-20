package routes

import (
	"app/internal/controllers"
	"app/internal/middleware"
	"net/http"
)

type Routes struct {
	Mux        *http.ServeMux
	Controller *controllers.Controller
	Middleware *middleware.Middleware
}

func New(controller *controllers.Controller, middleware *middleware.Middleware) *Routes {
	return &Routes{
		Controller: controller,
		Middleware: middleware,
		Mux:        http.NewServeMux(),
	}
}
func (r *Routes) Setup() {
	r.Mux.HandleFunc("/testSpeed", r.Controller.TestSpeed)
	r.Mux.HandleFunc("/checkAuth", r.Middleware.Auth(r.Controller.TestSpeed))
}
