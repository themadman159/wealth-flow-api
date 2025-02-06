package transactionrepo

import (
	"go-api/pkg/model"
	"go-api/types"

	"gorm.io/gorm"
)

type ITransactionRepository interface {
	Create(username string, req types.TransactionRequest) error
	GetAll(username string) ([]model.Transaction, error)
	Delete(username string, id int) error
}

type TransactionRepository struct {
	Database *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{
		Database: db,
	}
}
