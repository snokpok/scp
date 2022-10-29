package schema

import "github.com/golang-jwt/jwt"

type UserClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
