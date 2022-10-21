package startup

import (
	"log"

	"github.com/snokpok/scp-go/src/schema"
	"github.com/snokpok/scp-go/src/utils"
)

func SetupDB(dbcs *schema.DbClients) {
	// setup mongodb
	completionChan := make(chan string)
	go func() {
		mdb, err := utils.ConnectMongoDBSetup()
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
