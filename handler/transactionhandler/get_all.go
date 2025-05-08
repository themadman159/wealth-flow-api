package transactionhandler

import (
	"go-api/utils/jwtutil"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *TransactionHandler) GetAll(c *fiber.Ctx) error {

	username, err := jwtutil.GetUsernameFromToken(c)
	if err != nil {
		return h.Response.Unauthorized(c, "Unauthorized")
	}

	year := c.Query("year")
	if year == "" {
		return h.Response.BadRequest(c, "Invalid year")
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return h.Response.InternalServer(c, "Failed to convert int")
	}

	month := c.Query("month")
	if month == "" {
		return h.Response.BadRequest(c, "Invalid month")
	}
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return h.Response.InternalServer(c, "Failed to convert int")
	}

	searchByType := c.Query("search_type")
	searchByCategory := c.Query("search_category")

	transactions, err := h.TransactionService.GetAll(username, searchByType, searchByCategory, yearInt, monthInt)
	if err != nil {
		return h.Response.InternalServer(c, err.Error())
	}

	return h.Response.Success(c, "", transactions)
}
