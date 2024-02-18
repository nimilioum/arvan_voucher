package services

import (
	"arvan_voucher/core/models/user"
	"gorm.io/gorm"
)

func AddUser(tx *gorm.DB, user *user.User) (err error) {

	res := tx.Create(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil

}

func GetUser(tx *gorm.DB, phone string) (u *user.User, err error) {
	u = &user.User{Phone: phone}
	res := tx.First(&u)
	if res.Error != nil {
		return nil, res.Error
	}
	return u, nil
}
