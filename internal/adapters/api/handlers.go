package api

import (
	"app/internal/domain/user"
	"fmt"
	"net/http"
)

type Handlers struct {
	UserService user.AnyUserService
}

func NewHandlers(userService user.AnyUserService) *Handlers {
	return &Handlers{
		UserService: userService,
	}
}

func (c *Handlers) TestSpeed(w http.ResponseWriter, r *http.Request) {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//fmt.Fprintf(w, "Index %s\nURL %s", text, r.URL.String())
	fmt.Fprintf(w, "TestSpeed executed successfully")
}
