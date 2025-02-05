package userhandler

import (
	"go-api/types"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	var req types.UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Response.BadRequest(c, "Invalid request payload")
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return h.Response.BadRequest(c, err.Error())
	}

	if err := h.UserService.CreateUser(req); err != nil {
		return h.Response.InternalServer(c, err.Error())
	}

	return h.Response.Create(c, "Create User Successfully", nil)
}
