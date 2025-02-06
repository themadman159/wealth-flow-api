package loginroute

import (
	"go-api/handler/loginhandler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoginRoute(api fiber.Router, db *gorm.DB) {
	loginhandler := loginhandler.NewLoginHandler(db)

	api.Group("login").
		Post("/", loginhandler.LoginByUsernamePassword)
}
