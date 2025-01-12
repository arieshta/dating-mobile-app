package feeds

import (
	"context"
	"errors"
	"fmt"

	"github.com/arieshta/dating-mobile-app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Repository interface {
		GetByUserId(userId string) (*model.Feed, error)
		UpdateOne(filter bson.M, update bson.M) *model.Feed
		Create(feed *model.Feed) (*model.Feed, error)
	}

	Repo struct {
		holder *model.SharedHolder
		mongoCollection *mongo.Collection
	}
)

func (r *Repo) GetByUserId(userId string) (feed *model.Feed, err error) {
	fmt.Println(userId)
	err = r.mongoCollection.FindOne(context.Background(), bson.M{"user_id": userId}).Decode(&feed)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func (r *Repo) UpdateOne(filter bson.M, update bson.M) (feed *model.Feed) {
	err := r.mongoCollection.FindOneAndUpdate(context.Background(), filter, update).Decode(&feed)
	if err != nil {
		return nil
	}
	return feed
}

func (r *Repo) Create(feed *model.Feed) (*model.Feed, error) {
	result, err := r.mongoCollection.InsertOne(context.Background(), feed)
	if err != nil {
		return nil, err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		feed.ID = oid
	} else {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}
	return feed, nil
}

func NewRepository(holder *model.SharedHolder) (Repository, error) {
    if holder.Mongo == nil || holder.Config == nil {
        return nil, errors.New("holder is not properly initialized")
    }
	mongoCollection := holder.Mongo.Database(holder.Config.MongoDB).Collection("feed")
	return &Repo{holder, mongoCollection}, nil
}