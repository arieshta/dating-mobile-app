package users

import (
	"errors"

	"github.com/arieshta/dating-mobile-app/handler/auth"
	"github.com/arieshta/dating-mobile-app/model/dto"
)
func (s *ServiceImpl) LoginService(payload *dto.LoginBody) (string, error) {
	// Check if user exists
	user, err := s.userRepo.GetByUsername(payload.Username)
	if err != nil {
		return "", errors.New("wrong email or password")
	}

	// Compare password
	if auth.CheckPasswordHash(user.Password, payload.Password) {
		return "", errors.New("wrong email or password")
	}

	token, err := auth.CreateToken(user.ID.Hex())
	if err != nil {
		return "", err
	}
	return token, nil
}
