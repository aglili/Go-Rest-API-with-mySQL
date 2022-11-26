package config

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)


var db *gorm.DB

// Function to connect to mysql server
func Connect()  {
	d,err := gorm.Open("mysql","root:Selorm2001@tcp(127.0.0.1:3306)/trialdb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d	
}

func GetDB() *gorm.DB {
	return db
	
}