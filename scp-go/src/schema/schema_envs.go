package schema

type ExposedEnvironmentVariables struct {
	Port                int
	MongoDBClusterURI   string
	SecretJWT           string
	SpotifyClientID     string
	SpotifyClientSecret string
	RedisPassword       string
	RedisHost           string
	DeployMode          string
}
