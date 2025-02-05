package usersvc

import (
	"go-api/types"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) CreateUser(req types.UserCreateRequest) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := s.UserRepository.CreateUser(req.Username, string(hashedPassword), req.Email); err != nil {
		return err
	}

	return nil
}
