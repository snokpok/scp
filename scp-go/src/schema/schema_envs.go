package schema

type ExposedEnvironmentVariables struct {
	MongoDBClusterURI   string
	SecretJWT           string
	SpotifyClientID     string
	SpotifyClientSecret string
	RedisPassword       string
	RedisHost           string
	DeployMode          string
}
