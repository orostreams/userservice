package main

import (
	_ "net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/orostreams/userservice/config"
	"github.com/orostreams/userservice/controllers"
	"github.com/orostreams/userservice/utils"
)

type server struct {
	echo *echo.Echo
}

func newServer() *server {
	return &server{
		echo: echo.New(),
	}
}

func startServer(port string) {
	//controllers
	var (
		usersController = controllers.NewUserController()
	)
	server := newServer()
	server.echo.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Oro-Streams API"})
	})

	apiRoutes := server.echo.Group("/api/v1")
	apiRoutes.GET("/users", usersController.Index)
	apiRoutes.GET("/users/:id", usersController.GetById)
	apiRoutes.POST("/users", usersController.Create)
	apiRoutes.PUT("/users/:id", usersController.Update)
	apiRoutes.DELETE("/users/:id", usersController.Delete)

	server.echo.Logger.Fatal(server.echo.Start(port))
}

func main() {
	var err error
	//load configuration
	config.LoadEnvironmentVariables()
	//setup database connection
	if utils.ActiveConnection, err = utils.NewDatabaseConnector("mysql", os.Getenv("DATABASEPATH")); err != nil {
		log.Fatal("Cannot Connect to database" + err.Error())
	}

	//run migrations
	utils.RunMigrations()
	startServer(os.Getenv("SERVERPORT"))
}
