package services

import (
	"arvan_voucher/core/models/payment"
	"arvan_voucher/core/models/user"
	"errors"
	"gorm.io/gorm"
)

func CreateWallet(tx *gorm.DB, u *user.User) (err error) {
	w := payment.Wallet{
		Amount: 0,
		User:   u,
	}
	res := tx.Create(&w)
	err = res.Error

	return err
}

func Deposit(tx *gorm.DB, u *user.User, amount uint) (err error) {
	w := &payment.Wallet{User: u}
	res := tx.Find(&w)

	if res.Error != nil {
		return res.Error
	}

	w.Amount = w.Amount + amount
	tr := &payment.Transaction{
		User:   u,
		Type:   payment.Deposit,
		Amount: amount,
	}

	res = tx.Create(&tr)
	if res.Error != nil {
		return res.Error
	}

	res = tx.Save(&w)
	if res.Error != nil {
		return res.Error
	}

	return
}

func Withdraw(tx *gorm.DB, u *user.User, amount uint) (err error) {
	w := &payment.Wallet{User: u}
	res := tx.Find(&w)

	if res.Error != nil {
		return res.Error
	}

	if !w.HasEnoughBalance(amount) {
		return errors.New("not enough balance")
	}
	w.Amount = w.Amount - amount
	tr := &payment.Transaction{
		User:   u,
		Type:   payment.Withdraw,
		Amount: amount,
	}

	res = tx.Create(&tr)
	if res.Error != nil {
		return res.Error
	}

	res = tx.Save(&w)
	if res.Error != nil {
		return res.Error
	}

	return
}

func GetBalance(tx *gorm.DB, u *user.User) (w *payment.Wallet, err error) {
	w = new(payment.Wallet)
	w.User = u
	err = tx.First(&w).Error

	return
}

func GetTransactions(tx *gorm.DB, u *user.User) (transactions []*payment.Transaction, err error) {

	res := tx.Where(&payment.Transaction{User: u}).Find(&transactions)
	err = res.Error

	return
}
