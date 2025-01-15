package premium

import (
	"github.com/arieshta/dating-mobile-app/internal/repositories/premium"
	"github.com/arieshta/dating-mobile-app/model"
	"github.com/arieshta/dating-mobile-app/model/dto"
)

type (
	Service interface {
		HasPremium(userId string) bool
		Purchase(userId string, payload *dto.PurchaseBody) (*model.Premium, error)
	}

	ServiceImpl struct {
		sharedHolder *model.SharedHolder
		repo    premium.Repository
	}
)

func NewService(sharedHolder *model.SharedHolder) (*ServiceImpl, error) {
	premiumRepo, err := premium.NewRepository(sharedHolder)
	if err != nil {
		return nil, err
	}
	return &ServiceImpl{sharedHolder, premiumRepo}, nil
}
