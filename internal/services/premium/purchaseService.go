package premium

import (
	"errors"
	"time"

	"github.com/arieshta/dating-mobile-app/model"
	"github.com/arieshta/dating-mobile-app/model/dto"
)

func (s *ServiceImpl) Purchase(userId string, payload *dto.PurchaseBody) (*model.Premium, error) {
	existing, err := s.repo.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	today := time.Now()
	if existing != nil && !existing.PurchasedAt.AddDate(0, 0, existing.Duration).After(today) {
		return nil, errors.New("user still has premium")
	}
	premium := &model.Premium{
		UserId:      userId,
		Duration:    payload.Duration,
		PurchasedAt: time.Now(),
	}
	if existing != nil {
		premium.PurchaseNumber = existing.PurchaseNumber + 1
	} else {
		premium.PurchaseNumber = 1
	}

	_, err = s.repo.CreateOne(premium)
	if err != nil {
		return nil, err
	}
	return premium, nil
}
