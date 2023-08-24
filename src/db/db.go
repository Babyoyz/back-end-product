package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Error error

func ConectionMysql() (*gorm.DB, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	connection := os.Getenv("connection_db_mysql")
	db_name := os.Getenv("connection_db_mysql_name")
	db_charset := os.Getenv("connection_db_mysql_charset")

	dsn := fmt.Sprintf("root:root@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", connection, db_name, db_charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db, err

}
