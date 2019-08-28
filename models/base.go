package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("luizalabs")
	password := os.Getenv("root")
	dbName := os.Getenv("luizalabs")
	dbHost := os.Getenv("127.0.0.1:3306")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", username, password, dbHost, dbName)
	fmt.Println(dbUri)

	conn, err := gorm.Open("mysql", "luizalabs:root@tcp(127.0.0.1:3306)/luizalabs?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Customer{}, &CustomerList{}, &FavoritesProduct{})
}

func GetDB() *gorm.DB {
	return db
}
