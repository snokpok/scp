package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"
	"github.com/snokpok/scp-go/src/repositories/user"
	"github.com/snokpok/scp-go/src/schema"
	"github.com/snokpok/scp-go/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCurrentUser(c *gin.Context, dbcs *schema.DbClients) (*schema.User, int, error) {
	// get all user info from db with secret key
	email := c.GetString(string(schema.ContextMeClaim))
	users, err := user.FindUsers(dbcs.Mdb, bson.M{"email": email})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if len(*users) < 1 {
		return nil, 404, errors.New("user not found")
	}
	if len(*users) > 1 {
		return nil, 500, errors.New("incorrect number of user instances")
	}
	return &(*users)[0], http.StatusOK, nil
}

func GetUserById(id string, dbcs *schema.DbClients) (*schema.User, int, error) {
	user, err := user.FindOneUser(dbcs.Mdb, bson.M{"id": id})
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	return user, http.StatusAccepted, nil
}

type CreateUserResponse struct {
	User  schema.User `json:"username,omitempty"`
	Token string      `json:"token,omitempty"`
}

func CreateUser(c *gin.Context, dbcs *schema.DbClients) (*CreateUserResponse, int, error) {
	var userData schema.User
	var token string
	if err := c.ShouldBindWith(&userData, binding.JSON); err != nil {
		return nil, 400, err
	}
	userData.SecretKey = uuid.NewV4().String()
	insertCtx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	// try to create the user
	res, err := dbcs.Mdb.Database("main").Collection("users").InsertOne(insertCtx, userData)
	if err != nil {
		// if it's not already created => some other error to handle
		if !mongo.IsDuplicateKeyError(err) {
			return &CreateUserResponse{}, 400, err
		}
		// if it's created then handle creation of app-domain access token
		user, err := user.FindOneUser(dbcs.Mdb, bson.M{"email": userData.Email})
		if err != nil {
			// this often happens due to timeout or just some other problem interacting with the db
			return &CreateUserResponse{}, 400, err
		}
		token, _ = utils.GenerateAccessToken(utils.AuthTokenProps{
			ID:       user.ID.Hex(),
			Email:    userData.Email,
			Username: userData.Username,
		})
		return &CreateUserResponse{
			User:  *user,
			Token: token,
		}, 200, errors.New("user already created")
	}

	// generate token here
	token, err = utils.GenerateAccessToken(utils.AuthTokenProps{
		ID:       res.InsertedID.(primitive.ObjectID).Hex(),
		Email:    userData.Email,
		Username: userData.Username,
	})
	if err != nil {
		return nil, 500, err
	}
	return &CreateUserResponse{
		User:  userData,
		Token: token,
	}, 200, nil
}
