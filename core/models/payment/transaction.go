package payment

import (
	"arvan_voucher/core/models/user"
	"database/sql/driver"
	"gorm.io/gorm"
)

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
)

func (r *TransactionType) Scan(value interface{}) error {
	*r = TransactionType(value.([]byte))
	return nil
}

func (r TransactionType) Value() (driver.Value, error) {
	return string(r), nil
}

type Transaction struct {
	gorm.Model
	User   *user.User
	UserID uint            `json:"user_id"`
	Type   TransactionType `gorm:"type:transaction_type" json:"type"`
	Amount uint            `json:"amount"`
}
