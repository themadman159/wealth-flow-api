package userrepo

import (
	"errors"
	model "go-api/pkg/model/users"

	"gorm.io/gorm"
)

func (r *UserRepository) CreateUser(username, password, email string) error {

	var userAlready *model.User
	result := r.Database.Where("email = ? OR username = ?", email, username).
		First(&userAlready)

	if result.Error == nil {
		return errors.New("User Already Exists")
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	user := &model.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	if err := r.Database.Model(&model.User{}).Create(&user).Error; err != nil {
		return err
	}

	return nil
}
