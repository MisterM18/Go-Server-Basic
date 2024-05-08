package handlers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	DB *gorm.DB
}

type Customer struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

func (h *CustomerHandler) Initialize() {
	db, err := gorm.Open(mysql.Open("webservice:P@ssw0rd@tcp(127.0.0.1:3306)/db_webservice?charset=utf8&parseTime=True"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Customer{})

	h.DB = db
}
