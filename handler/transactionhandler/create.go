package transactionhandler

import (
	"go-api/types"
	"go-api/utils/jwtutil"
	"go-api/utils/validateutil"

	"github.com/gofiber/fiber/v2"
)

func (h *TransactionHandler) Create(c *fiber.Ctx) error {

	username, err := jwtutil.GetUsernameFromToken(c)
	if err != nil {
		return h.Response.Unauthorized(c, "Unauthorized")
	}

	var req types.TransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Response.BadRequest(c, "Invalid request payload")
	}

	err = validateutil.Validate(req)
	if err != nil {
		return h.Response.BadRequest(c, err.Error())
	}

	err = h.TransactionService.Create(username, req)
	if err != nil {
		return h.Response.InternalServer(c, err.Error())
	}

	return h.Response.Create(c, "Create transaction successfully", nil)
}
