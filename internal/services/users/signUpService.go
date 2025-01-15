package users

import (
	"errors"

	"github.com/arieshta/dating-mobile-app/handler/auth"
	"github.com/arieshta/dating-mobile-app/model"
	"github.com/arieshta/dating-mobile-app/model/dto"
)

func (s *ServiceImpl) SignUpService(payload *dto.SignUpBody) (*model.User, error) {
	_, err := s.userRepo.GetByUsername(payload.Username)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	_, err = s.userRepo.GetByEmail(payload.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Username: payload.Username,
		Password: hashedPassword,
		Email:    payload.Email,
		Fullname: payload.Fullname,
		Gender:   payload.Gender,
		Picture:  payload.Picture,
	}

	_, err = s.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
