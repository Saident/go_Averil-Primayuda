package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("localhost"),
		os.Getenv("root"),
		os.Getenv(""),
		os.Getenv("go_clean"),
		os.Getenv("3306"),
	)

	var err error

	dbConn, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	e := echo.New()

	NewRoute(e, dbConn)

	e.Logger.Fatal(e.Start(":" + os.Getenv("8000")))
}
