package routes

import (
	"arvan_voucher/api/handlers"
	"github.com/labstack/echo/v4"
)

func addUserRoutes(router *echo.Group) {
	router.POST("/", handlers.AddUserHandler)
}
