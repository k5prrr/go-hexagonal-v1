package controllers

import (
	"net/http"
	"log"
	"fmt"
)

var testSpeedI int64 = 0
func TestSpeed(w http.ResponseWriter, r *http.Request) {
	testSpeedI++
	fmt.Fprintf(w, "Index %d\nURL %s", testSpeedI, r.URL.String())
	log.Println("Еще одно сообщение.123")

	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(users)
}
