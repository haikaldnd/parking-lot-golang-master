package configs

import (
	"fmt"

	"gitlab.mapan.io/playground/parking-lot-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USER = "root"
const DB_PASS = ""
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"
const DB_NAME = "parking"

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Parking{})

}
