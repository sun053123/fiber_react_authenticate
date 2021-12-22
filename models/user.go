package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model        //เข้าไปแก้ใน ตัว gorm.Model ให้ส่งแค่ id เป็น lowwercase
	Username   string `json:"username"`
	Email      string `gorm:"typevarchar(100);unique" json:"email"`
	Password   string `validate:"required,password,min=6,max=32" json:"-"`
	Age        string `json:"-"`
	Occupation string `json:"-"`
	Posts      []Post `json:"-"`
}
