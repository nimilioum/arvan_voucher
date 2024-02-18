package payment

import (
	"arvan_voucher/core/models/user"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	UserID uint `gorm:"uniqueIndex" json:"user_id"`
	User   *user.User
	Amount uint
}

func (w *Wallet) HasEnoughBalance(amount uint) bool {
	return w.Amount >= amount
}
