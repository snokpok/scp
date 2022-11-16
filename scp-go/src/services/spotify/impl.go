package spotify

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snokpok/scp-go/src/repositories/spotify"
	"github.com/snokpok/scp-go/src/repositories/user"
	"github.com/snokpok/scp-go/src/schema"
	"github.com/snokpok/scp-go/src/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFromSpotifyCurrentlyPlaying(c *gin.Context, dbcs *schema.DbClients) (*map[string]interface{}, int, error) {

	email := c.GetString(string(schema.ContextMeClaim))

	userFound, err := user.FindOneUser(dbcs.Mdb, bson.M{"email": email})
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	result, code, err := GetSCPForUser(userFound, dbcs)
	return result, code, err
}

func GetSCPForUser(user *schema.User, dbcs *schema.DbClients) (*map[string]interface{}, int, error) {
	resultScp, _ := spotify.RequestSCPFromSpotify(user.AccessToken)

	if resultScp["error"] != nil {
		// request refreshed access token from spotify
		utils.LOUT.Printf("--Refreshing new access token for user %s from Spotify API--\n", user.ID.String())
		newTkn, err := spotify.RequestNewAccessTokenFromSpotify(user.RefreshToken)
		if err != nil {
			utils.LERR.Printf("Couldn't refresh access token: %s\n", err)
			return nil, http.StatusFailedDependency, err
		}

		// update the newly issued access token from spotify
		updateCmd := bson.M{
			"$set": bson.M{"access_token": newTkn},
		}
		filter := bson.M{"id": user.ID}
		user.UpdateOneUser(dbcs.Mdb, filter, updateCmd)

		utils.LOUT.Println("Fetching new SCP results...")
		// fetch the new results
		resultScp, err = spotify.RequestSCPFromSpotify(newTkn)
		if err != nil {
			utils.LERR.Printf("Couldn't fetch new SCP results with new token: %s\n", err)
			return nil, http.StatusFailedDependency, err
		}
	}

	return &resultScp, http.StatusAccepted, nil
}
