package middlewares

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/snokpok/scp-go/src/repositories/user"
	"github.com/snokpok/scp-go/src/schema"
	"github.com/snokpok/scp-go/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MwAuthorizeCurrentUser(mdb *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("--Authorizing user--")
		token, err := utils.HelperGetTokenValidateHeader(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := utils.DecodeAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		// create child context with the student claims
		c.Set("user", user.Email)
		c.Set("claims", user)
		c.Next()
	}
}

func MwAuthorizeBasicHeaderRefresh(mdb *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("--With refresh header Basic--")
		secret, err := utils.HelperGetTokenValidateHeader(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := user.FindOneUser(mdb, bson.M{"secret_key": secret})
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx := context.WithValue(c.Request.Context(), schema.ContextMeClaim, user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
