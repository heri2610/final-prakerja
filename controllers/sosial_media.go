package controllers

import (
	"github.com/heri2610/final-prakerja/config"
	"github.com/heri2610/final-prakerja/models"
	"github.com/labstack/echo/v4"
)

func AddSosialMediaController(c echo.Context) error {
	var sosialMedia models.SosialMedia
	c.Bind(&sosialMedia)

	err := config.DB.Create(&sosialMedia).Error
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
		Data:    sosialMedia,
	})
}

func GetDetailSosialMediaController(c echo.Context) error {
	// id := c.Param("id")
	var sosialMedia models.SosialMedia = models.SosialMedia{}
	return c.JSON(200, sosialMedia)
}

func GetSosialMediaController(c echo.Context) error {
	var sosialMedia []models.SosialMedia
	err := config.DB.Find(&sosialMedia).Error
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
		Data:    sosialMedia,
	})
}
