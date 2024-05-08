package main

import (
	"mongkol/api/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/Hello", handlers.HelloWorld)

	e.Logger.Fatal(e.Start(":8080"))
}
