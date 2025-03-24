package Server

import (
	"fmt"
	//"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type ServerConfig struct {
	Address string `json:"address"`
	Release bool   `json:"release"`
	Ssl     bool   `json:"ssl"`
	Crt     string `json:"crt"`
	Key     string `json:"key"`
}

type Server struct {
	Engine *gin.Engine
	config *ServerConfig
}

func New(config *ServerConfig) *Server {
	// Конфиг по умолчанию
	if config == nil {
		config = &ServerConfig{
			Address: ":8080",
		}
	}

	// Устанавливаем режим Release
	if config.Release {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	// Отдаём элемент сервера
	return &Server{
		Engine: gin.Default(),
		config: config,
	}
}

func (server *Server) Run() error {
	fmt.Println(fmt.Sprintf("Server start %s", server.config.Address))

	var err error
	if server.config.Ssl {
		err = server.Engine.RunTLS(
			server.config.Address,
			server.config.Crt,
			server.config.Key,
		)

	} else {
		err = server.Engine.Run(server.config.Address)

	}
	if err != nil {
		return err
	}

	return nil
}

/*
err = http.ListenAndServeTLS(
	server.config.Address,
	server.config.Crt,
	server.config.Key,
	nil,
)
err = http.ListenAndServe(
	server.config.Address,
	nil,
)
*/
