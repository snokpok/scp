package startup

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/snokpok/scp-go/src/schema"
)

var Psv schema.ExposedEnvironmentVariables

// load in server environments and configure various server settings accordingly
// e.g DEPLOY_MODE.
// If was not parsed before then parse it again
func LoadServerEnv(file string) schema.ExposedEnvironmentVariables {
	// load in envfile
	err := godotenv.Load(file)
	if err != nil {
		log.Fatal(err)
	}

	if Psv != nil {
		return Psv
	}

	Psv := schema.ExposedEnvironmentVariables{
		MongoDBClusterURI:   os.Getenv("MONGODB_CLUSTER_URI"),
		SecretJWT:           os.Getenv("SECRET_JWT"),
		SpotifyClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedisPassword:       os.Getenv("RDB_DEFAULT_PASSWORD"),
		RedisHost:           os.Getenv("REDIS_HOST"),
		DeployMode:          os.Getenv("DEPLOY_MODE"),
	}
	return Psv
}
