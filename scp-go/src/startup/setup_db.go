package startup

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/snokpok/scp-go/src/schema"
	"github.com/snokpok/scp-go/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDB(dbcs *schema.DbClients) {
	// setup mongodb
	completionChan := make(chan string)
	go func() {
		mdb, err := ConnectMongoDBSetup()
		if err != nil {
			log.Fatal(err)
		}
		// set them in db clients
		dbcs.Mdb = mdb
		// collections
		completionChan <- "mdb"
	}()
	<-completionChan
}

func ConnectMongoDBSetup() (*mongo.Client, error) {
	// starting up the database client with timeout of 5s
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	uri := utils.LoadServerEnv().MongoDBClusterURI
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
