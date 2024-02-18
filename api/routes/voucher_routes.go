package routes

import (
	"arvan_voucher/api/handlers"
	"github.com/labstack/echo/v4"
)

func addVoucherRoutes(router *echo.Group) {
	router.POST("/", handlers.AddVoucherHandler)
	router.POST("/use/", handlers.UseVoucherHandler)
}
