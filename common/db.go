package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-swagger-sample/models"
)

func InitializeDb() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:12541254@tcp(127.0.0.1:3306)/installer?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Consumer{})

	return db, err
}
