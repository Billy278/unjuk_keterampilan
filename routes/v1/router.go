package router

import (
	"unjuk_keterampilan/controllers"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.GET("/products", controllers.ShowAllProduct)
	e.POST("/products", controllers.AddProduct)
	e.GET("/products/:id", controllers.FindById)
	e.PUT("/products/:id", controllers.Updateproduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)
}
