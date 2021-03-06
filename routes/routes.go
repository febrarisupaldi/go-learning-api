package routes

import (
	"net/http"

	"github.com/febrarisupaldi/go-learning-api/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/customer", controllers.GetAllCustomer)
	e.POST("/customer", controllers.AddCustomers)
	e.POST("/test", controllers.BindJSON)

	return e
}
