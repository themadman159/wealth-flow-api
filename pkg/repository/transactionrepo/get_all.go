package transactionrepo

import (
	"errors"
	"go-api/pkg/model"

	"gorm.io/gorm"
)

func (r *TransactionRepository) GetAll(username string) ([]model.Transaction, error) {

	var user model.User
	if err := r.Database.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	var transaction []model.Transaction
	if err := r.Database.Where("user_id = ?", user.ID).Find(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}
