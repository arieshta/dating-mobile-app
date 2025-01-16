package matches

import (
	"fmt"
	"slices"

	"github.com/arieshta/dating-mobile-app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ServiceImpl) LikeOneService(userId, matchId string) (*model.Likes, error) {
	_, err := s.userRepo.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	likes, err := s.likeRepo.GetAllByUserId(userId)
	if err == mongo.ErrNoDocuments {
		newLikes := &model.Likes{
			UserID: userId,
			LikeIDs: []primitive.ObjectID{},
		}
		likes, err = s.likeRepo.Create(newLikes)
		fmt.Println("new likes created", likes)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	newLikeId, err := primitive.ObjectIDFromHex(matchId)
	if err != nil {
		return nil, err
	}
	if slices.Contains(likes.LikeIDs, newLikeId) {
		return likes, nil
	}
	likes = s.likeRepo.UpdateOne(bson.M{"_id": likes.ID}, bson.M{"$push": bson.M{"like_ids": newLikeId}})
	return likes, nil
}