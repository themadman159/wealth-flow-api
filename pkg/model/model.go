package model

type DefaultBy struct {
	CreateBy int `gorm:"column:created_by"`
	UpdateBy int `gorm:"column:updated_by"`
	DeleteBy int `gorm:"column:deleted_by"`
}
