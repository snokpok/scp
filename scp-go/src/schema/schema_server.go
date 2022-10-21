package schema

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbClients struct {
	Mdb *mongo.Client
	Rdb *redis.Client
}