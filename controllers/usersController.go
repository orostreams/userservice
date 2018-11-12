package controllers

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/orostreams/userservice/services"
)

type UserController struct {
	service services.ServiceInterface
}

func NewUserController() UserController {
	return UserController{
		service: services.NewUserService(),
	}
}

func (c UserController) Index(context echo.Context) error {
	users, err := c.service.GetAll()
	if err != nil {
		return context.JSON(500, map[string]interface{}{
			"Error": "Internal Server Error",
			"Logs":  err.Error(),
		})
	}

	return context.JSON(200, map[string]interface{}{
		"Message": "Success",
		"Data":    users,
	})
}

func (c UserController) Create(context echo.Context) error {
	input := map[string]interface{}{}
	//TODO Validate input
	if err := context.Bind(&input); err != nil {
		return context.JSON(500, map[string]interface{}{
			"Error": "Internal Server Error",
			"Logs":  err.Error(),
		})
	}
	user, err := c.service.Create(input)
	if err != nil {
		return context.JSON(500, map[string]interface{}{
			"Error": "Internal Server Error",
			"Logs":  err.Error(),
		})
	}
	return context.JSON(200, map[string]interface{}{
		"Message": "Success",
		"Data":    user,
	})
}

func (c UserController) GetById(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return context.JSON(500, map[string]interface{}{
			"Error": "Internal Server Error",
			"Logs":  err.Error(),
		})
	}
	user, err := c.service.GetByID(id)
	if err != nil {
		return context.JSON(500, map[string]interface{}{
			"Error": "Internal Server Error",
			"Logs":  err.Error(),
		})
	}

	return context.JSON(200, map[string]interface{}{
		"Message": "Success",
		"Data":    user,
	})
}

func (c UserController) Update(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return context.JSON(500, map[string]interface{}{
			"Error": "Internal Server Error",
			"Logs":  err.Error(),
		})
	}

	input := map[string]interface{}{}
	if err := context.Bind(&input); err != nil {
		return context.JSON(500, map[string]interface{}{
			"Error": "Internal Server Error",
			"Logs":  err.Error(),
		})
	}
	user, err := c.service.Update(id, input)
	if err != nil {
		return context.JSON(500, map[string]interface{}{
			"Error": "Internal Server Error",
			"Logs":  err.Error(),
		})
	}
	return context.JSON(200, map[string]interface{}{
		"Message": "Success",
		"Data":    user,
	})
}

func (c UserController) Delete(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return context.JSON(500, map[string]interface{}{
			"Error": "Internal Server Error",
			"Logs":  err.Error(),
		})
	}
	if err := c.service.Delete(id); err != nil {
		return context.JSON(500, map[string]interface{}{
			"Error": "Internal Server Error",
			"Logs":  err.Error(),
		})
	}
	return context.JSON(200, map[string]interface{}{
		"Message": "Success",
	})
}
