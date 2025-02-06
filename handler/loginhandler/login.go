package loginhandler

import (
	"go-api/pkg/service/loginsvc"
	"go-api/utils/responseutil"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ILoginHandler interface {
	LoginByUsernamePassword(c *fiber.Ctx) error
}

type LoginHandler struct {
	LoginService loginsvc.ILoginService
	Response     responseutil.ResponseUtil
}

func NewLoginHandler(db *gorm.DB) ILoginHandler {
	return &LoginHandler{
		LoginService: loginsvc.NewLoginService(db),
	}
}
