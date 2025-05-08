package transactionsvc

import (
	"go-api/pkg/repository/transactionrepo"
	"go-api/types"

	"gorm.io/gorm"
)

type ITransactionService interface {
	Create(username string, req types.TransactionRequest) error
	GetAll(username, searchType, searchCategory string, year, month int) (*types.TransactionGetAllResponse, error)
	Delete(username string, id int) error
}

type TransactionService struct {
	TransactionRepository transactionrepo.ITransactionRepository
}

func NewTransactionService(db *gorm.DB) ITransactionService {
	return &TransactionService{
		TransactionRepository: transactionrepo.NewTransactionRepository(db),
	}
}
