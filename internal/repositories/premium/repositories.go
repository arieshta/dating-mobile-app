package premium

import (
	"context"
	"errors"
	"time"

	"github.com/arieshta/dating-mobile-app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Repository interface {
		CreateOne(premium *model.Premium) (*model.Premium, error)
		GetByUserId(userId string) (*model.Premium, error)
		UpdateOne(userId string, premium bool) *model.Premium
	}

	Repo struct {
		holder          *model.SharedHolder
		mongoCollection *mongo.Collection
	}
)

func (r *Repo) GetByUserId(userId string) (premium *model.Premium, err error) {
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	err = r.mongoCollection.FindOne(context.Background(), bson.M{"user_id": userIdObj}).Decode(&premium)
	if err != nil {
		return nil, err
	}

	return premium, nil
}

func (r *Repo) UpdateOne(userId string, premium bool) (prem *model.Premium) {
	update := bson.M{
		"$set":         bson.M{"premium": premium},
		"$currentDate": bson.M{"updated_at": true},
	}

	err := r.mongoCollection.FindOneAndUpdate(context.Background(), bson.M{"user_id": userId}, update).Decode(&prem)
	if err != nil {
		return nil
	}

	return prem
}

func (r *Repo) CreateOne(premium *model.Premium) (prem *model.Premium, err error) {
	premium.CreatedAt = time.Now()
	premium.UpdatedAt = premium.CreatedAt

	res, err := r.mongoCollection.InsertOne(context.Background(), premium)
	if err != nil {
		return nil, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		premium.ID = oid
	} else {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	return premium, nil
}

func NewRepository(holder *model.SharedHolder) (Repository, error) {
	if holder.Mongo == nil || holder.Config == nil {
		return nil, errors.New("holder is not properly initialized")
	}

	mongoCollection := holder.Mongo.Database(holder.Config.MongoDB).Collection("premium")

	return &Repo{holder, mongoCollection}, nil
}
