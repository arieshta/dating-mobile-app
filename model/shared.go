package model

import (
	"context"
	"log"

	"github.com/arieshta/dating-mobile-app/internal/config"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	SharedHolder struct {
		Config *config.Config
		Echo   *echo.Echo
		Mongo  *mongo.Client
	}
)

func (s *SharedHolder) Init() {
	s.Config = config.LoadEnv()
	s.Echo = echo.New()
	s.Mongo = initMongo(s.Config)

}

func initMongo(config *config.Config) *mongo.Client {
	opts := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("1",err)
		return nil
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal("2",err)
		return nil
	}
	return client
}