package startup

import (
	"github.com/gin-gonic/gin"
	"github.com/snokpok/scp-go/src/controllers"
	mws "github.com/snokpok/scp-go/src/middlewares"
	"github.com/snokpok/scp-go/src/schema"
)

func CreateRouter(dbcs *schema.DbClients) *gin.Engine {
	r := gin.Default()

	r.Use(mws.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	r.POST("/user", controllers.CreateUser(dbcs))
	r.GET("/user/:id", controllers.GetUserSCP(dbcs))
	r.GET("/me", mws.MwAuthorizeCurrentUser(dbcs.Mdb), controllers.GetMe(dbcs))
	r.GET("/scp", mws.MwAuthorizeCurrentUser(dbcs.Mdb), controllers.GetSCP(dbcs))

	return r
}
