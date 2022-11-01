package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/snokpok/scp-go/src/schema"
	"github.com/snokpok/scp-go/src/startup"
	"github.com/snokpok/scp-go/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbcs *schema.DbClients = &schema.DbClients{}

var (
	UserCol *mongo.Collection
)

func main() {
	lerr := log.New(os.Stderr, log.Prefix(), 0)

	// load in envfile if there is any
	if _, err := os.Open(".env"); err == nil {
		utils.LoadServerEnv(".env")
	}
	// setup mongodb
	startup.SetupDB(dbcs)
	// router setup
	r := startup.CreateRouter(dbcs)

	port := utils.LoadServerEnv().Port
	// invalid port parsed from LoadServerEnv (<0 or NaN)
	if port == -1 {
		lerr.Println("Invalid port (must be positive integer) -- defaulting to 4000")
		port = 4000 // default to this
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
