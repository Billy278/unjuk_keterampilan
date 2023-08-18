package config

import (
	"fmt"
	"os"
	"unjuk_keterampilan/models"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
}

func InitDatabase() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_ROOT"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	initMigration()
}

func initMigration() {

	DB.AutoMigrate(&models.Product{})
}
