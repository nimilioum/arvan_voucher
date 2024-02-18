package routes

import (
	"arvan_voucher/api/handlers"
	"github.com/labstack/echo/v4"
)

func addPaymentRoutes(router *echo.Group) {
	router.GET("/:phone/", handlers.GetBalanceHandler)
	router.POST("/deposit/", handlers.DepositHandler)
	router.POST("/withdraw/", handlers.WithdrawHandler)
	router.GET("/transactions/:phone/", handlers.GetTransactionsHandler)
}
