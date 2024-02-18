package handlers

import (
	"arvan_voucher/core/models/user"
	"arvan_voucher/services"
	"arvan_voucher/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func AddUserHandler(c echo.Context) (err error) {

	u := new(user.User)
	if err = c.Bind(u); err != nil {
		fmt.Printf(err.Error())
		return c.JSON(http.StatusBadRequest, &utils.ErrorMessage{Message: "Invalid user"})
	}

	db := c.Get("DB").(*gorm.DB)
	tx := db.Begin()

	err = services.AddUser(tx, u)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = services.CreateWallet(tx, u)
	if err != nil {
		tx.Rollback()
		return err
	}

	res := tx.Commit()
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	return c.NoContent(http.StatusCreated)
}
