package matches

import "github.com/arieshta/dating-mobile-app/model"

func (s *ServiceImpl) PassOneService(userId, matchId string) (*model.Feed, error) {
	feed, err := s.feedRepo.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	return feed, nil
}