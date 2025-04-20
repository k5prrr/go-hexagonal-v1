package controllers

import (
	"app/internal/services"
	"fmt"
	"log"
	"net/http"
)

type Controller struct {
	Services *services.Services
}

func New(services *services.Services) *Controller {
	//fmt.Printf("%v", services)
	return &Controller{
		Services: services,
	}
}

func (c *Controller) TestSpeed(w http.ResponseWriter, r *http.Request) {
	// Теперь можно использовать c.service
	text, err := c.Services.TestSpeed()
	if err != nil {
		log.Printf("Error in testSpeed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Index %s\nURL %s", text, r.URL.String())
	//fmt.Fprintf(w, "TestSpeed executed successfully")

}
