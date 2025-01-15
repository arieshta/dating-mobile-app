package users

import (
	"github.com/arieshta/dating-mobile-app/model"
)

func (s *ServiceImpl) GetOneByUserIdService(userId string) (*model.User, error) {
	user, err := s.userRepo.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}