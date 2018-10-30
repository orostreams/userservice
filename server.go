package http

import (
	_ "net/http"

	"github.com/labstack/echo"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	return &Server{
		echo: echo.New(),
	}
}

func Start(port string) {
	server := NewServer()
	server.echo.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Oro-Streams API"})
	})

	server.echo.Logger.Fatal(server.echo.Start(port))
}
