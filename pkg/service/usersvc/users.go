package usersvc

import (
	"go-api/pkg/repository/userrepo"
	"go-api/types"

	"gorm.io/gorm"
)

type IUserService interface {
	CreateUser(req types.UserCreateRequest) error
}

type UserService struct {
	UserRepository userrepo.IUserRepository
}

func NewUserService(db *gorm.DB) IUserService {
	return &UserService{
		UserRepository: userrepo.NewUserRepository(db),
	}
}
