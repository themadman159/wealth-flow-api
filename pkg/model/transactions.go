package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	DefaultBy
	Type            string    `json:"type" gorm:"column:type;type:enum('INCOME','EXPENSE')"`
	Category        string    `json:"category" gorm:"column:category"`
	Amount          int64     `json:"amount" gorm:"column:amount"`
	Description     string    `json:"description" gorm:"column:description"`
	TransactionDate time.Time `json:"transaction_date" gorm:"column:transaction_date"`
	UserID          uint      `json:"user_id" gorm:"column:user_id"`
	User            User      `json:"user" gorm:"foreignKey:UserID"`
}
