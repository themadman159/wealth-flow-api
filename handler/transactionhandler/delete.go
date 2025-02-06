package transactionhandler

import (
	"go-api/utils/jwtutil"

	"github.com/gofiber/fiber/v2"
)

func (h *TransactionHandler) Delete(c *fiber.Ctx) error {

	username, err := jwtutil.GetUsernameFromToken(c)
	if err != nil {
		return h.Response.Unauthorized(c, "Unauthorized")
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return h.Response.BadRequest(c, "Invalid transaction ID")
	}

	h.TransactionService.Delete(username, id)

	return h.Response.Success(c, "Delete transaction success fully", nil)
}
