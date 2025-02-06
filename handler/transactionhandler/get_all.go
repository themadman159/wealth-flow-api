package transactionhandler

import (
	"go-api/utils/jwtutil"

	"github.com/gofiber/fiber/v2"
)

func (h *TransactionHandler) GetAll(c *fiber.Ctx) error {

	username, err := jwtutil.GetUsernameFromToken(c)
	if err != nil {
		return h.Response.Unauthorized(c, "Unauthorized")
	}

	transactions, err := h.TransactionService.GetAll(username)
	if err != nil {
		return h.Response.InternalServer(c, err.Error())
	}

	return h.Response.Success(c, "", transactions)
}
