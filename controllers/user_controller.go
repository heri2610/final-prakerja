package controllers

import (
	"github.com/heri2610/final-prakerja/config"
	"github.com/heri2610/final-prakerja/models"
	"github.com/labstack/echo/v4"
)

func AddUsersController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	err := config.DB.Create(&user)

	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed create in database",
			Data:    nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Successfully created",
		Data:    user,
	})
}

func GetDetailUsersController(c echo.Context) error {
	// id := c.Param("id")
	var user models.User = models.User{}
	return c.JSON(200, user)
}

func GetUsersController(c echo.Context) error {
	var users []models.User
	err := config.DB.Find(users)
	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}
	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Successfully get data",
		Data:    users,
	})
}
