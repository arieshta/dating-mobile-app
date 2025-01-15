package users

import (
	"github.com/arieshta/dating-mobile-app/model"
	"github.com/arieshta/dating-mobile-app/model/dto"
)

func (s *ServiceImpl) UpdateService(userId string, payload *dto.UpdateBody) (*model.User, error) {
	return s.userRepo.Update(userId, payload)
}