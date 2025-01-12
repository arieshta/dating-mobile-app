package premium

import (
	"github.com/arieshta/dating-mobile-app/internal/repositories/users"
	"github.com/arieshta/dating-mobile-app/model"
	"github.com/arieshta/dating-mobile-app/model/dto"
)

type (
	Service interface {
		Purchase(payload *dto.LoginBody) (string, error)
	}

	ServiceImpl struct {
		sharedHolder *model.SharedHolder
		userRepo    users.Repository
	}
)

func NewService(sharedHolder *model.SharedHolder) (*ServiceImpl, error) {
	userRepo, err := users.NewRepository(sharedHolder)
	if err != nil {
		return nil, err
	}
	return &ServiceImpl{sharedHolder, userRepo}, nil
}
