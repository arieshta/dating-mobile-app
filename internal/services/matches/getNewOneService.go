package matches

import (
	"fmt"

	"github.com/arieshta/dating-mobile-app/model"
	"github.com/arieshta/dating-mobile-app/model/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ServiceImpl) GetNewOneFeedService(userId string) (*model.Feed, error) {
	// Check if user exists
	user, err := s.userRepo.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	// Check if user has feed
	feed, err := s.feedRepo.GetByUserId(userId)
	if err == mongo.ErrNoDocuments {
		newFeed := &model.Feed{
			UserID: userId,
			Feeds: []primitive.ObjectID{},
		}
		feed, err = s.feedRepo.Create(newFeed)
		fmt.Println("new feed created", feed)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	if (s.premiumService.HasPremium(userId)) {
		dailyLimit := 10
		if len(feed.Feeds) >= dailyLimit {
			return nil, fmt.Errorf("daily limit reached")
		}
	}
	var genderFilter string
	if user.Gender == "male" {
		genderFilter = "female"
	} else {
		genderFilter = "male"
	}
	newFeed, err := s.userRepo.GetRandomOneWithFilter(dto.GetRandomFilter{Gender: genderFilter, ExcludedIds: feed.Feeds})
	if err != nil {
		return nil, err
	}
	feed = s.feedRepo.UpdateOne(bson.M{"_id": feed.ID}, bson.M{"$push": bson.M{"feeds": newFeed.ID}})
	// feed.Feeds = append(feed.Feeds, newFeed.ID)
	return feed, nil
}