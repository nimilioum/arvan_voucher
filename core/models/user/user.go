package user

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Phone     string `gorm:"uniqueIndex" json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
