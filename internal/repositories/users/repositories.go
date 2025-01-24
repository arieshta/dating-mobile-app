package users

import (
	"context"
	"errors"
	"time"

	"github.com/arieshta/dating-mobile-app/model"
	"github.com/arieshta/dating-mobile-app/model/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Repository interface {
		GetByUserId(userId string) (*model.User, error)
		GetByUsername(username string) (*model.User, error)
		GetByEmail(email string) (*model.User, error)
		GetRandomOneWithFilter(filter dto.GetRandomFilter) (*model.User, error)
		Create(user *model.User) (*model.User, error)
		Update(userId string, payload *dto.UpdateBody) (*model.User, error)
		Delete(userId string) error
	}

	Repo struct {
		holder          *model.SharedHolder
		mongoCollection *mongo.Collection
	}
)

func (r *Repo) GetByUserId(userId string) (user *model.User, err error) {
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	err = r.mongoCollection.FindOne(context.Background(), bson.M{"_id": userIdObj}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repo) GetByUsername(username string) (user *model.User, err error) {
	err = r.mongoCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repo) GetByEmail(email string) (user *model.User, err error) {
	err = r.mongoCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repo) GetRandomOneWithFilter(filter dto.GetRandomFilter) (user *model.User, err error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "gender", Value: filter.Gender},
            {Key: "_id", Value: bson.D{{Key: "$nin", Value: filter.ExcludedIds}}},
		}}},
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
	}

	cursor, err := r.mongoCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no user found")
	}

	return user, nil
}

func (r *Repo) Create(user *model.User) (*model.User, error) {
    user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	result, err := r.mongoCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	} else {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	return user, nil
}

func (r *Repo) Update(userId string, payload *dto.UpdateBody) (user *model.User, err error) {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New("failed to convert userId to ObjectID")
	}

	opts := options.FindOneAndUpdateOptions{}
	opts.SetReturnDocument(options.After)

    update := bson.M{
        "$set": payload,
        "$currentDate": bson.M{"updatedAt": true},
    }

	err = r.mongoCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": oid}, update, &opts).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repo) Delete(userId string) error {
	_, err := r.mongoCollection.DeleteOne(context.Background(), bson.M{"_id": userId})
	if err != nil {
		return err
	}

	return nil
}

func NewRepository(holder *model.SharedHolder) (Repository, error) {
	if holder.Mongo == nil || holder.Config == nil {
		return nil, errors.New("holder is not properly initialized")
	}

	mongoCollection := holder.Mongo.Database(holder.Config.MongoDB).Collection("users")

	return &Repo{holder, mongoCollection}, nil
}
