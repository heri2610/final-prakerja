package main

import (
	"fmt"
	"os"

	"github.com/heri2610/final-prakerja/config"
	"github.com/heri2610/final-prakerja/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	config.InitDB()
	e := echo.New()
	routes.InitRoute(e)
	port := os.Getenv("PORT")
	e.Start(fmt.Sprintf(":%s", port))
}
