package utils

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDBSetup() (*mongo.Client, error) {
	// starting up the database client with timeout of 5s
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	uri := LoadServerEnv().MongoDBClusterURI
	clientConfigs := options.Client().ApplyURI(uri)
	mdb, err := mongo.Connect(ctx, clientConfigs)
	if err != nil {
		return nil, err
	}
	log.Println("successfully connected to database!")

	return mdb, nil
}

func CreateIndexesMDB(mdb *mongo.Client) {
	// creating indices in mongodb
	ivModels := []mongo.IndexModel{
		{
			Keys:    bson.D{primitive.E{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("email_unique"),
		},
	}
	UserCol := mdb.Database("main").Collection("users")
	opts := options.CreateIndexes().SetMaxTime(5 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	createManyIndexChan := make(chan int, 1)
	var names []string
	go func() {
		namesCreated, err := UserCol.Indexes().CreateMany(ctx, ivModels, opts)
		names = namesCreated
		if err != nil {
			log.Println("failed in creating indexes " + strings.Join(namesCreated, ", "))
		}
		createManyIndexChan <- 1
	}()
	defer cancel()
	select {
	case <-ctx.Done():
		log.Println("timed out creating indexes")
	case <-createManyIndexChan:
		log.Println("successfully created indexes: ", strings.Join(names, ", "))
	}
}

func ConnectSetupRedis() (*redis.Client, error) {
	// connecting with redis client
	envs := LoadServerEnv()
	rdb := redis.NewClient(&redis.Options{
		Addr:     envs.RedisHost,
		Password: envs.RedisPassword,
	})

	ctxPingRDB, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	statCmd := rdb.Ping(ctxPingRDB)
	if statCmd.Err() != nil {
		return nil, statCmd.Err()
	}
	log.Println("successfully connected to redis cluster!")
	return rdb, nil
}
