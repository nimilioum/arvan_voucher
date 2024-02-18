package main

import (
	"arvan_voucher/api/routes"
	"arvan_voucher/config"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	err := godotenv.Load(".env")
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	db, err := config.InitDB()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}
	e.Use(config.DBMiddleware(db))

	err = config.Migrate(db)
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	apiRouter := e.Group("/api")
	routes.AddApiRoutes(apiRouter)

	e.Logger.Fatal(e.Start("localhost:5000"))

}
