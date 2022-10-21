package repositories

import (
	"context"
	"time"

	"github.com/snokpok/scp-go/src/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUsers(mdb *mongo.Client, filter bson.M) (*[]schema.User, error) {
	// looks up the secret key in the database to verify it belongs to someone
	lookupCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	users := []schema.User{}
	cursor, err := mdb.Database("main").Collection("users").Find(lookupCtx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		user := schema.User{}
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func FindOneUser(mdb *mongo.Client, filter bson.M) (*schema.User, error) {
	// looks up the secret key in the database to verify it belongs to someone
	lookupCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	user := schema.User{}
	err := mdb.Database("main").Collection("users").FindOne(lookupCtx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateOneUser(mdb *mongo.Client, filter bson.M, update bson.M) (*schema.User, error) {
	ctxTO, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := mdb.Database("main").Collection("users").UpdateOne(ctxTO, filter, update)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
