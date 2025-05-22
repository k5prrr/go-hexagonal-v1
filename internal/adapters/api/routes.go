package api

import (
	"net/http"
)

type Routes struct {
	Mux      *http.ServeMux
	Handlers *Handlers
}

func NewRoutes(handlers *Handlers) *Routes {
	return &Routes{
		Mux:      http.NewServeMux(),
		Handlers: handlers,
	}
}
func (r *Routes) Setup() {
	r.Mux.HandleFunc("/testSpeed", r.Handlers.TestSpeed)
	//r.Mux.HandleFunc("/checkAuth", r.Middleware.Auth(r.Controller.TestSpeed))
}
