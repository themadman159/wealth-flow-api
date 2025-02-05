package examplehandler

import (
	"go-api/pkg/service/examplesvc"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IExampleHandler interface {
	Example(c *fiber.Ctx) error
}

type ExampleHandler struct {
	ExampleService examplesvc.IExampleService
}

func NewExampleHandler(dbconn *gorm.DB) IExampleHandler {
	return &ExampleHandler{
		ExampleService: examplesvc.NewExampleServiced(dbconn),
	}
}

func (h *ExampleHandler) Example(c *fiber.Ctx) error {
	example := h.ExampleService.Example()
	return c.Status(fiber.StatusOK).JSON(example)
}
