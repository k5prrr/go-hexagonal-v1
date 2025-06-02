package main

import (
	"app/pkg/env"
	"fmt"
)

func main() {
	env := env.New("")
	fmt.Println(env.Get("POSTGRES_USER", "us1"))

}
