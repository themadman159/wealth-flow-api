package userroute

import (
	"go-api/handler/userhandler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoute(api fiber.Router, db *gorm.DB) {

	userhandler := userhandler.NewUserHandler(db)
	api.Group("/users").
		Post("/", userhandler.CreateUser)
}
