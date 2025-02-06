package transactionroute

import (
	"go-api/handler/transactionhandler"
	"go-api/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TransactionRoute(api fiber.Router, db *gorm.DB) {

	transactionhandler := transactionhandler.NewTransactionHandler(db)
	api.Group("/transactions").Use(middleware.AuthMiddleware()).
		Get("/", transactionhandler.GetAll).
		Post("/", transactionhandler.Create).
		Delete("/:id", transactionhandler.Delete)
}
