package example

import (
	"go-api/handler/examplehandler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ExampleRoute(api fiber.Router, db *gorm.DB) {

	examplehandler := examplehandler.NewExampleHandler(db)
	api.Group("/example").
		Get("/", examplehandler.Example)
}
