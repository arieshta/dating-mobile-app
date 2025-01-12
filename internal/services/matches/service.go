package matches

import (
	"github.com/arieshta/dating-mobile-app/internal/repositories/feeds"
	"github.com/arieshta/dating-mobile-app/internal/repositories/likes"
	"github.com/arieshta/dating-mobile-app/internal/repositories/users"
	"github.com/arieshta/dating-mobile-app/model"
)

type (
	Service interface {
		// GetOne(userId string) (*model.Likes, error)
		GetLikesService(userId string) (*model.Likes, error)
		GetNewOneFeedService(userId string) (*model.Feed, error)
		LikeOneService(userId, matchId string) (*model.Likes, error)
	}

	ServiceImpl struct {
		sharedHolder *model.SharedHolder
		feedRepo     feeds.Repository
		likeRepo     likes.Repository
		userRepo     users.Repository
	}
)

func NewService(sharedHolder *model.SharedHolder) (*ServiceImpl, error) {
	feedRepo, err := feeds.NewRepository(sharedHolder)
	if err != nil {
		return nil, err
	}
	likeRepo, err := likes.NewRepository(sharedHolder)
	if err != nil {
		return nil, err
	}
	userRepo, err := users.NewRepository(sharedHolder)
	if err != nil {
		return nil, err
	}
	return &ServiceImpl{sharedHolder, feedRepo, likeRepo, userRepo}, nil
}
