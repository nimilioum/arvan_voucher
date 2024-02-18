package handlers

import (
	"arvan_voucher/core/models/voucher"
	"arvan_voucher/core/schemas"
	"arvan_voucher/services"
	"arvan_voucher/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func AddVoucherHandler(c echo.Context) (err error) {
	v := new(voucher.Voucher)
	if err = c.Bind(v); err != nil {
		fmt.Printf(err.Error())
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: "Invalid request"})
	}

	db := c.Get("DB").(*gorm.DB)
	tx := db.Begin()

	err = services.AddVoucher(tx, v)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: err.Error()})
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func UseVoucherHandler(c echo.Context) (err error) {
	sc := new(schemas.UseVoucherRequest)
	if err = c.Bind(sc); err != nil {
		fmt.Printf(err.Error())
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: "Invalid request"})
	}

	db := c.Get("DB").(*gorm.DB)
	tx := db.Begin()

	u, err := services.GetUser(tx, sc.Phone)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: err.Error()})
	}

	v, err := services.GetVoucher(tx, sc.Code)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: err.Error()})
	}

	err = services.UseVoucher(tx, u, v)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: err.Error()})
	}

	err = services.Deposit(tx, u, v.Amount)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: err.Error()})
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}
