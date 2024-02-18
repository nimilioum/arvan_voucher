package handlers

import (
	"arvan_voucher/core/schemas"
	"arvan_voucher/services"
	"arvan_voucher/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func DepositHandler(c echo.Context) (err error) {

	sc := new(schemas.DepositWithdrawRequest)
	if err = c.Bind(sc); err != nil {
		fmt.Printf(err.Error())
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: "Invalid request"})
	}

	db := c.Get("DB").(*gorm.DB)
	tx := db.Begin()

	u, err := services.GetUser(tx, sc.Phone)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = services.Deposit(tx, u, sc.Amount)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: err.Error()})
	}

	res := tx.Commit()
	if res.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: res.Error.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func WithdrawHandler(c echo.Context) (err error) {

	sc := new(schemas.DepositWithdrawRequest)
	if err = c.Bind(sc); err != nil {
		fmt.Printf(err.Error())
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: "Invalid request"})
	}

	db := c.Get("DB").(*gorm.DB)
	tx := db.Begin()

	u, err := services.GetUser(tx, sc.Phone)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: err.Error()})
	}

	err = services.Withdraw(tx, u, sc.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: err.Error()})
	}

	res := tx.Commit()
	if res.Error != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: res.Error.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func GetTransactionsHandler(c echo.Context) (err error) {
	phone := c.Param("phone")
	db := c.Get("DB").(*gorm.DB)
	tx := db

	u, err := services.GetUser(tx, phone)
	if err != nil {
		return err
	}

	transactions, err := services.GetTransactions(tx, u)
	if err != nil {
		return err
	}

	res := tx.Commit()
	if res.Error != nil {
		return res.Error
	}

	return c.JSON(http.StatusOK, &transactions)
}

func GetBalanceHandler(c echo.Context) (err error) {
	phone := c.Param("phone")
	db := c.Get("DB").(*gorm.DB)
	tx := db

	u, err := services.GetUser(tx, phone)
	if err != nil {
		return err
	}

	w, err := services.GetBalance(tx, u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &w)
}
