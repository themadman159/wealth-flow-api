package model

import (
	"go-api/pkg/model"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	model.DefaultBy
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Email    string `json:"email" gorm:"column:email"`
}
