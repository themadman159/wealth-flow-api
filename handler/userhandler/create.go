package userhandler

import (
	"go-api/types"
	"go-api/utils/validateutil"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	var req types.UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Response.BadRequest(c, "Invalid request payload")
	}

	err := validateutil.Validate(req)
	if err != nil {
		return h.Response.BadRequest(c, err.Error())
	}

	if err := h.UserService.CreateUser(req); err != nil {
		return h.Response.InternalServer(c, err.Error())
	}

	return h.Response.Create(c, "Create User Successfully", nil)
}
