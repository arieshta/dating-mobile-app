package matches

import (
	"github.com/arieshta/dating-mobile-app/model"
)

func (s *ServiceImpl) GetLikesService(userId string) (*model.Likes, error) {
	// Check if user exists
	_, err := s.userRepo.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	// Check if user has likes
	likes, err := s.likeRepo.GetAllByUserId(userId)
	if err != nil {
		return nil, err
	}
	return likes, nil
}