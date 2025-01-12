package users

import (
	"errors"
	"fmt"

	"github.com/arieshta/dating-mobile-app/model"
	"github.com/arieshta/dating-mobile-app/model/dto"
)

func (s *ServiceImpl) SignUpService(payload *dto.SignUpBody) (*model.User, error) {
	// Check if user exists
	_, err := s.userRepo.GetByUsername(payload.Username)
	if err == nil {
		return nil, errors.New("user already exists")
	}
	_, err = s.userRepo.GetByEmail(payload.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}
	// Create user
	user := model.User{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
		Fullname: payload.Fullname,
		Gender:   payload.Gender,
		Picture:  payload.Picture,
	}
	_, err = s.userRepo.Create(&user)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
