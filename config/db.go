package config

import (
	"arvan_voucher/core/models/payment"
	"arvan_voucher/core/models/user"
	_ "embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

//go:embed create_enums.sql
var enumsQuery string

func InitDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_NAME"), os.Getenv("POSTGRES_PORT"))
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func migrateEnums(db *gorm.DB) (err error) {

	return db.Exec(enumsQuery).Error
}

func Migrate(db *gorm.DB) (err error) {
	err = migrateEnums(db)
	if err != nil {
		return err
	}

	return db.AutoMigrate(
		&user.User{},
		&payment.Wallet{},
		&payment.Transaction{},
	)
}

func DBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("DB", db)
			return next(c)
		}
	}
}
