package database

import (
	model "go-api/pkg/model"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {

	models := []interface{}{
		&model.User{},
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}
