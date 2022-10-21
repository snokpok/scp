package schema

type UserBody struct {
    Username string `json:"username,omitempty" bson:"username,omitempty" binding:"required"`
    Email string `json:"email,omitempty" bson:"email,omitempty" binding:"required"`
    SpotifyId string `json:"spotify_id,omitempty" bson:"spotify_id,omitempty" binding:"required"`
    AccessToken string `json:"access_token,omitempty" bson:"access_token,omitempty" binding:"required"`
    RefreshToken string `json:"refresh_token,omitempty" bson:"refresh_token,omitempty" binding:"required"`
}
