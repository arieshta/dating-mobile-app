package likes

import (
	"context"
	"errors"
	"time"

	"github.com/arieshta/dating-mobile-app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Repository interface {
		GetAllByUserId(userId string) (*model.Likes, error)
		Create(likes *model.Likes) (*model.Likes, error)
		UpdateOne(bson.M, bson.M) (*model.Likes)
	}

	Repo struct {
		holder *model.SharedHolder
		mongoCollection *mongo.Collection
	}
)

func (r *Repo) GetAllByUserId(userId string) (likes *model.Likes, err error) {
	err = r.mongoCollection.FindOne(context.Background(), bson.M{"user_id": userId}).Decode(&likes)
	if err != nil {
		return nil, err
	}

	return likes, nil
}

// func (r *Repo) GetByEmail(email string) (likes *model.Likes, err error) {
// 	err = r.mongoCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&likes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return likes, nil
// }

func (r *Repo) Create(likes *model.Likes) (*model.Likes, error) {
	likes.CreatedAt = time.Now()
	likes.UpdatedAt = likes.CreatedAt

	result, err := r.mongoCollection.InsertOne(context.Background(), likes)
	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		likes.ID = oid
	} else {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	return likes, nil
}

func (r *Repo) UpdateOne(filter, update bson.M) (likes *model.Likes) {
	opts := options.FindOneAndUpdateOptions{}
	opts.SetReturnDocument(options.After)

	update["updated_at"] = time.Now()

	err := r.mongoCollection.FindOneAndUpdate(context.Background(), filter, update, &opts).Decode(&likes)
	if err != nil {
		return nil
	}

	return likes
}

func NewRepository(holder *model.SharedHolder) (Repository, error) {
    if holder.Mongo == nil || holder.Config == nil {
        return nil, errors.New("holder is not properly initialized")
    }

	mongoCollection := holder.Mongo.Database(holder.Config.MongoDB).Collection("likes")

	return &Repo{holder, mongoCollection}, nil
}