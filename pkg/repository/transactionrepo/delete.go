package transactionrepo

import (
	"errors"
	"go-api/pkg/model"
	"time"

	"gorm.io/gorm"
)

func (r *TransactionRepository) Delete(username string, id int) error {

	var user model.User
	if err := r.Database.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if err := r.Database.Model(&model.Transaction{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"deleted_by": user.ID,
		}).Error; err != nil {
		return err
	}

	return nil
}
