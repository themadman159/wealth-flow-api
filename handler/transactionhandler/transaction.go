package transactionhandler

import (
	"go-api/pkg/service/transactionsvc"
	"go-api/utils/responseutil"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ITransactionHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type TransactionHandler struct {
	TransactionService transactionsvc.ITransactionService
	Response           responseutil.ResponseUtil
}

func NewTransactionHandler(db *gorm.DB) ITransactionHandler {
	return &TransactionHandler{
		TransactionService: transactionsvc.NewTransactionService(db),
	}
}
