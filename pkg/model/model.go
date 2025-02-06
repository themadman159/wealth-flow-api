package model

type DefaultBy struct {
	CreateBy uint `gorm:"column:created_by"`
	UpdateBy uint `gorm:"column:updated_by"`
	DeleteBy uint `gorm:"column:deleted_by"`
}
