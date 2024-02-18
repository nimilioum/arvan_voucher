package routes

import "github.com/labstack/echo/v4"

func AddApiRoutes(e *echo.Group) {
	userRouter := e.Group("/user")
	paymentRouter := e.Group("/payment")
	voucherRouter := e.Group("/voucher")

	addUserRoutes(userRouter)
	addPaymentRoutes(paymentRouter)
	addVoucherRoutes(voucherRouter)
}
