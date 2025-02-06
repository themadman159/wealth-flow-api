package routes

import (
	"go-api/routes/example"
	"go-api/routes/loginroute"
	"go-api/routes/transactionroute"
	"go-api/routes/userroute"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	example.ExampleRoute(api, db)
	userroute.UserRoute(api, db)
	loginroute.LoginRoute(api, db)
	transactionroute.TransactionRoute(api, db)

}
