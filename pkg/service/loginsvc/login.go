package loginsvc

import (
	"go-api/pkg/repository/loginrepo"
	"go-api/types"

	"gorm.io/gorm"
)

type ILoginService interface {
	LoginByUsernamePassword(req types.LoginByUsernamePasswordRequest) (*types.LoginResponse, error)
}

type LoginService struct {
	LoginRepository loginrepo.ILoginRepository
}

func NewLoginService(db *gorm.DB) ILoginService {
	return &LoginService{
		LoginRepository: loginrepo.NewLoginRepository(db),
	}
}
