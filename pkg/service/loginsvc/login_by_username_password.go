package loginsvc

import (
	"errors"
	"go-api/types"
	"go-api/utils/jwtutil"

	"golang.org/x/crypto/bcrypt"
)

func (s *LoginService) LoginByUsernamePassword(req types.LoginByUsernamePasswordRequest) (*types.LoginResponse, error) {

	user, err := s.LoginRepository.LoginByUsernamePassword(req.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	accessToken, err := jwtutil.CreateToken(user.Username, user.Email)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		AccessToken: accessToken, // Replace with actual JWT token generation
		Username:    user.Username,
		Email:       user.Email,
	}, nil
}
