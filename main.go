package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	MySQL = "MYSQL"
)

func main() {
	//App
	app := fiber.New()

	//DB
	var db *gorm.DB
	var err error

	if len(os.Getenv(MySQL)) > 0 {
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_URL")
		dbName := os.Getenv("DB_NAME")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=MySQL", user, pass, dbName,
			dbHost)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	}

	port := "1337"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if err != nil {
		log.Fatal(err)
	}

	_ = db

	if err = app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}

}
