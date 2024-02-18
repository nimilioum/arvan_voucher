package voucher

import (
	"arvan_voucher/core/models/user"
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	Code   string       `json:"code" gorm:"uniqueIndex"`
	Amount uint         `json:"amount"`
	Cap    uint         `json:"cap"`
	Count  uint         `json:"count"`
	Users  []*user.User `gorm:"many2many:voucher_users;"`
}
