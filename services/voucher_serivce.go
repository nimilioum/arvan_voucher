package services

import (
	"arvan_voucher/core/models/user"
	"arvan_voucher/core/models/voucher"
	"errors"
	"gorm.io/gorm"
)

func AddVoucher(tx *gorm.DB, v *voucher.Voucher) (err error) {
	err = tx.Create(&v).Error

	return err
}

func GetVoucher(tx *gorm.DB, code string) (v *voucher.Voucher, err error) {
	v = &voucher.Voucher{Code: code}
	err = tx.Find(&v).Error
	if err != nil {
		return nil, err
	}

	return
}

func UseVoucher(tx *gorm.DB, u *user.User, v *voucher.Voucher) (err error) {
	if v.Count >= v.Cap {
		return errors.New("Voucher  expired")
	}

	err = tx.Model(&v).Association("Users").Append(u)
	if err != nil {
		return err
	}

	v.Count++
	err = tx.Save(&v).Error

	return
}
