package controllers

import (
	"net/http"

	"github.com/febrarisupaldi/go-learning-api/models"
	"github.com/labstack/echo/v4"
)

func GetAllCustomer(c echo.Context) error {
	result, err := models.GetAllCustomer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
