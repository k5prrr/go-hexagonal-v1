package api

import (
	"net/http"
)

type Routes struct {
	Mux        *http.ServeMux
	Handlers   *Handlers
	Middleware *Middleware
}

func NewRoutes(handlers *Handlers, middleware *Middleware) *Routes {
	return &Routes{
		Mux:        http.NewServeMux(),
		Handlers:   handlers,
		Middleware: middleware,
	}
}
func (r *Routes) Setup() {
	//r.Mux.HandleFunc("/testSpeed", r.Handlers.TestSpeed)
	r.Mux.HandleFunc("/api/users", r.Middleware.CheckApiKey(r.Handlers.Users))
	//r.Mux.HandleFunc("/checkAuth", r.Middleware.Auth(r.Controller.TestSpeed))
}
