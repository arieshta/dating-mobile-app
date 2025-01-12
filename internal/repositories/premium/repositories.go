package premium

import (
	"context"
	"errors"

	"github.com/arieshta/dating-mobile-app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Repository interface {
		GetByUserId(userId string) (*model.Premium, error)
		UpdateOne(userId string, premium bool) (*model.Premium)
	}

	Repo struct {
		holder *model.SharedHolder
		mongoCollection *mongo.Collection
	}
)

func (r *Repo) GetByUserId(userId string) (premium *model.Premium, err error) {
	err = r.mongoCollection.FindOne(context.Background(), bson.M{"user_id": userId}).Decode(&premium)
	if err != nil {
		return nil, err
	}
	return premium, nil
}

func (r *Repo) UpdateOne(userId string, premium bool) (prem *model.Premium) {
	err := r.mongoCollection.FindOneAndUpdate(context.Background(), bson.M{"user_id": userId}, bson.M{"premium": premium}).Decode(&prem)
	if err != nil {
		return nil
	}
	return prem
}

func NewRepository(holder *model.SharedHolder) (Repository, error) {
    if holder.Mongo == nil || holder.Config == nil {
        return nil, errors.New("holder is not properly initialized")
    }
	mongoCollection := holder.Mongo.Database(holder.Config.MongoDB).Collection("premium")
	return &Repo{holder, mongoCollection}, nil
}
