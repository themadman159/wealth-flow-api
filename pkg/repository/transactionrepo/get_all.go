package transactionrepo

import (
	"errors"
	"go-api/pkg/model"

	"gorm.io/gorm"
)

func (r *TransactionRepository) GetAll(username, searchType, searchCategory string, year, month int) ([]model.Transaction, error) {

	var user model.User
	if err := r.Database.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	var transaction []model.Transaction

	tx := r.Database.Model(&model.Transaction{}).Where("user_id = ?", user.ID)

	if year != 0 && month != 0 {
		tx = tx.Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", month, year)
	}

	if searchType != "" {
		tx = tx.Where("type = ?", searchType)
	}

	if searchCategory != "" {
		tx = tx.Where("category = ?", searchCategory)
	}

	if err := tx.Order("transaction_date DESC").Find(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}
