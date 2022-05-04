package middleware_redis_simple_cache

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodb *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.3.11:27017"))
	if err != nil {
		panic(err)
	}
	mongodb = client

	if err := mongodb.Ping(ctx, nil); err != nil {
		panic(err.Error())
	}
}

func GetDB() *mongo.Client {
	return mongodb
}
