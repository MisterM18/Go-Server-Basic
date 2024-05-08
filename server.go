import (
	"github.com/labstack/echo/v4"
	"github.com/your/package/handlers"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	h := handlers.CustomerHandler{}
	h.Initialize()

	e.GET("/customers", h.GetAllCustomer)
	e.POST("/customers", h.SaveCustomer)
	e.GET("/customers/:id", h.GetCustomer)
	e.PUT("/customers/:id", h.UpdateCustomer)
	e.DELETE("/customers/:id", h.DeleteCustomer)
	e.Logger.Fatal(e.Start(":8080"))
}
