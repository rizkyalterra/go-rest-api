package routes

import (
	"project/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUsersController)

	return e
}
