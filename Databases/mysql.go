package Databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Mysql() {
	var err error
	DB, err = gorm.Open("mysql", "root:Mysql112604..@(1.15.144.204:3306)/course?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("mysql err:", err)
	}
	if DB.Error != nil {
		fmt.Println("DB error:", DB.Error)
	}
}
