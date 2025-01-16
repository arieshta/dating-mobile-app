package premium

import (
	"fmt"
	"time"
)

func (s *ServiceImpl) HasPremium(userId string) bool {
	premium, err := s.repo.GetByUserId(userId)
	fmt.Println(premium, userId)
	if err != nil || premium == nil {
		return false
	}
	today := time.Now()
	return !premium.PurchasedAt.AddDate(0, 0, premium.Duration).After(today)
}
