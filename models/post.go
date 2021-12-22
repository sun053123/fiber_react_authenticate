package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Body     string `gorm:"typevarchar(300)"`
	URI      string
	Comments []Comment
	UserId   int
}

type Comment struct {
	gorm.Model
	Comment string
	PostId  string
	UserId  int
}
