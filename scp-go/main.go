package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	schema "github.com/snokpok/scp-go/src/schema"
	"github.com/snokpok/scp-go/src/startup"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbcs *schema.DbClients = &schema.DbClients{}

var (
	UserCol *mongo.Collection
)

func main() {
	// load in envfile
	startup.LoadServerEnv(".env")
	// setup mongodb
	startup.SetupDB(dbcs)
	// router setup
	r := startup.CreateRouter(dbcs)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4000"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
