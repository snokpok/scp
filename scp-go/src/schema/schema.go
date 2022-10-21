package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Username string `json:"username,omitempty" bson:"username,omitempty"`
    Email string `json:"email,omitempty" bson:"email,omitempty"`
    SpotifyId string `json:"spotify_id,omitempty" bson:"spotify_id,omitempty"`
    SecretKey string `json:"secret_key,omitempty" bson:"secret_key,omitempty"`
    AccessToken string `json:"access_token,omitempty" bson:"access_token,omitempty"`
    RefreshToken string `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
}
