package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/snokpok/scp-go/src/schema"
	"github.com/snokpok/scp-go/src/services"
)

func CreateUser(dbcs *schema.DbClients) gin.HandlerFunc {
	// create user, store (username, email, spotify_id, access_token, refresh_token)
	// if conflict user then don't do anything
	// new user will have id->email entry in redis
	return func(c *gin.Context) {
		res, code, err := services.CreateUser(c, dbcs)
		if err != nil {
			c.AbortWithStatusJSON(code, gin.H{
				"error": err.Error(),
				"data":  *res,
			})
			return
		}

		c.JSON(code, gin.H{
			"data": *res,
		})
	}
}

func GetMe(dbcs *schema.DbClients) gin.HandlerFunc {
	// get all user info from db with secret key
	return func(c *gin.Context) {
		user, code, err := services.GetCurrentUser(c, dbcs)
		if err != nil {
			c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
			return
		}
		c.JSON(code, *user)
	}
}

func GetSCP(dbcs *schema.DbClients) gin.HandlerFunc {
	// get the currently playing song for the user
	return func(c *gin.Context) {
		resultScp, code, err := services.GetFromSpotifyCurrentlyPlaying(c, dbcs)
		if err != nil {
			c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
			return
		}
		c.JSON(code, *resultScp)
	}
}
