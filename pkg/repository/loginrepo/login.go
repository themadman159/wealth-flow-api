package loginrepo

import (
	"errors"
	"go-api/pkg/model"

	"gorm.io/gorm"
)

type ILoginRepository interface {
	LoginByUsernamePassword(username string) (*model.User, error)
}

type LoginRepository struct {
	Database *gorm.DB
}

func NewLoginRepository(db *gorm.DB) ILoginRepository {
	return &LoginRepository{
		Database: db,
	}
}

func (r *LoginRepository) LoginByUsernamePassword(username string) (*model.User, error) {

	var user model.User
	if err := r.Database.Model(&model.User{}).
		Where("username = ?", username).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	return &user, nil
}
