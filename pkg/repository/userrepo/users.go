package userrepo

import "gorm.io/gorm"

type IUserRepository interface {
	CreateUser(name, password, email string) error
}

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		Database: db,
	}
}
