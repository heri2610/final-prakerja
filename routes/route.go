package routes

import (
	"github.com/heri2610/final-prakerja/controllers"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetDetailUsersController)
	e.POST("/users", controllers.AddUsersController)
	// ini untuk sosial media
	e.GET("/sosial-media", controllers.GetSosialMediaController)
	e.GET("/sosial-media/:id", controllers.GetDetailSosialMediaController)
	e.POST("/sosial-media", controllers.AddSosialMediaController)
}
