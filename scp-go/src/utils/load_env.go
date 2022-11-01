package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/snokpok/scp-go/src/schema"
)

var psv *schema.ExposedEnvironmentVariables

// load in server environments and configure various server settings accordingly
// e.g DEPLOY_MODE.
// If was not parsed before then parse it again
func LoadServerEnv(files ...string) *schema.ExposedEnvironmentVariables {
	// load in envfile if there is any specified
	// if can't parse file then panics
    if len(files) != 0 && err := godotenv.Load(files...); err != nil {
		log.Panic(err)
	}

	if psv != nil {
		return psv
	}

	// handle invalid port (<0 & NaN)
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil || port < 0 {
		port = -1
	}
	psv := schema.ExposedEnvironmentVariables{
		Port:                port,
		MongoDBClusterURI:   os.Getenv("MONGODB_CLUSTER_URI"),
		SecretJWT:           os.Getenv("SECRET_JWT"),
		SpotifyClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedisPassword:       os.Getenv("RDB_DEFAULT_PASSWORD"),
		RedisHost:           os.Getenv("REDIS_HOST"),
		DeployMode:          os.Getenv("DEPLOY_MODE"),
	}
	return &psv
}
