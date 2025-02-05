package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DefaultBy
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Email    string `json:"email" gorm:"column:email"`
}
