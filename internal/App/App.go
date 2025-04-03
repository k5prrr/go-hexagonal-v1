package App

import (
	"fmt"
	"App/pkg/Config"
)

func Main() {
	fmt.Println("hw7")
	fmt.Println(Config.Bool("debug"))
}
