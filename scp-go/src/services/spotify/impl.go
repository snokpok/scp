package spotify

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snokpok/scp-go/src/repositories/spotify"
	"github.com/snokpok/scp-go/src/repositories/user"
	"github.com/snokpok/scp-go/src/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFromSpotifyCurrentlyPlaying(c *gin.Context, dbcs *schema.DbClients) (*map[string]interface{}, int, error) {

	email := c.GetString(string(schema.ContextMeClaim))

	userFound, err := user.FindOneUser(dbcs.Mdb, bson.M{"email": email})
	if err != nil {
		return nil, 404, err
	}

	resultScp, _ := spotify.RequestSCPFromSpotify(userFound.AccessToken)

	if resultScp["error"] != nil {
		// request refreshed access token from spotify
		log.Println("--refreshing new access token from spotify--")
		newTkn, err := spotify.RequestNewAccessTokenFromSpotify(userFound.RefreshToken)
		if err != nil {
			return nil, http.StatusFailedDependency, err
		}

		// update the newly issued access token from spotify
		updateCmd := bson.M{
			"$set": bson.M{"access_token": newTkn},
		}
		dbcs.Mdb.Database("main").Collection("users").FindOneAndUpdate(
			context.Background(),
			bson.M{"email": email},
			updateCmd,
		)

		// fetch the new results
		resultScp, err = spotify.RequestSCPFromSpotify(newTkn)
		if err != nil {
			return nil, http.StatusFailedDependency, err
		}
	}

	return &resultScp, 200, nil
}
