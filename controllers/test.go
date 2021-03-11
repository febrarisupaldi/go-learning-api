package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func BindJSON(c echo.Context) error {
	var u []User
	// u := new([]User)
	var u2 []User

	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for index, user := range u {
		if index != 0 {
			u2 = append(u2, user)
		}
	}
	return c.JSON(http.StatusOK, u2)
}
