package transactionrepo

import (
	"errors"
	"go-api/pkg/model"
	"go-api/types"

	"gorm.io/gorm"
)

func (r *TransactionRepository) Create(username string, req types.TransactionRequest) error {

	var user model.User
	if err := r.Database.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	transaction := &model.Transaction{
		Description:     req.Description,
		Amount:          req.Amount,
		Type:            req.Type,
		Category:        req.Category,
		TransactionDate: req.TransactionDate,
		UserID:          user.ID,
		DefaultBy: model.DefaultBy{
			CreateBy: user.ID,
			UpdateBy: user.ID,
			DeleteBy: 0,
		},
	}

	if err := r.Database.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}
