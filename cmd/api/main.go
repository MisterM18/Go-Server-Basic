// cmd/main.go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/your/package/handlers"
)

func main() {
	e := echo.New()

	h := handlers.NewUserHandler() // สร้าง instance ของ UserHandler

	// กำหนด Route โดยใช้ UserHandler ที่ถูกสร้างขึ้น
	e.GET("/users", h.GetAllUser)
	e.POST("/users", h.SaveUser)
	e.GET("users/:id", h.GetUser)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
