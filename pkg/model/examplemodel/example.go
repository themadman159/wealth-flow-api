package model

import (
	"go-api/pkg/model"

	"gorm.io/gorm"
)

type ExampleModel struct {
	gorm.Model
	model.DefaultBy
	Message string `json:"message"`
}
