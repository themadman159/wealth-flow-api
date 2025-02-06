package example

import (
	"go-api/handler/examplehandler"
	"go-api/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ExampleRoute(api fiber.Router, db *gorm.DB) {

	examplehandler := examplehandler.NewExampleHandler(db)
	api.Group("/example").Use(middleware.AuthMiddleware()).
		Get("/", examplehandler.Example)
}
