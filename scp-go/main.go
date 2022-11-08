package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	env := utils.LoadServerEnv()

	// load in envfile if there is any
	if _, err := os.Open(".env"); err == nil {
		env = utils.LoadServerEnv(".env")
	}
	// setup mongodb
	startup.SetupDB(dbcs)
	// router setup
	r := startup.CreateRouter(dbcs)

	if env.DeployMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := env.Port
	// invalid port parsed from LoadServerEnv (<0 or NaN)
	if port == -1 {
		utils.LERR.Println("Invalid port (must be positive integer) -- defaulting to 4000")
		port = 4000 // default to this
	}

	utils.LOUT.Printf("Server listening on port %d", port)
	utils.LERR.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
