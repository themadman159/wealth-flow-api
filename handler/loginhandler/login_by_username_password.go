package loginhandler

import (
	"go-api/types"
	"go-api/utils/validateutil"

	"github.com/gofiber/fiber/v2"
)

func (h *LoginHandler) LoginByUsernamePassword(c *fiber.Ctx) error {

	var req types.LoginByUsernamePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Response.BadRequest(c, "Invalid request payload")
	}

	err := validateutil.Validate(req)
	if err != nil {
		return h.Response.BadRequest(c, err.Error())
	}

	user, err := h.LoginService.LoginByUsernamePassword(req)
	if err != nil {
		return h.Response.InternalServer(c, err.Error())
	}

	return h.Response.Success(c, "", user)
}
