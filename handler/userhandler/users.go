package userhandler

import (
	"go-api/pkg/service/usersvc"
	"go-api/utils/responseutil"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IUserHandler interface {
	CreateUser(c *fiber.Ctx) error
}

type UserHandler struct {
	UserService usersvc.IUserService
	Response    responseutil.ResponseUtil
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		UserService: usersvc.NewUserService(db),
	}
}
