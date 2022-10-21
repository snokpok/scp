package utils

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

type AuthTokenProps struct {
	ID       interface{}
	Username string
	Email    string
}

func HelperGetTokenValidateHeader(authHeader string) (string, error) {
	splitHeader := strings.Split(authHeader, " ")
	if len(splitHeader) < 2 {
		return "", errors.New("no header; unauthorized")
	}
	if splitHeader[0] != "Basic" && splitHeader[0] != "Bearer" {
		return "", errors.New("incorrect token scheme; must be basic or bearer")
	}
	secret := splitHeader[1]
	if secret == "" {
		return "", errors.New("unauthorized")
	}
	return secret, nil
}

func GenerateAccessToken(userData AuthTokenProps) (string, error) {
	// create the jwt token to authorize client to THIS server (not Spotify's)
	secretKey := []byte(os.Getenv("SECRET_JWT"))
	claims := UserClaim{
		userData.Username,
		userData.Email,
		jwt.StandardClaims{
			Issuer: os.Getenv("SPOTIFY_CLIENT_ID"),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func DecodeAccessToken(token string) (UserClaim, error) {
	// decoding the app auth token; returns empty UserClaim struct with err if there's an error
	claims := UserClaim{}
	tokenInfo, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if tokenInfo == nil || !tokenInfo.Valid {
		return claims, err
	}
	return claims, nil
}
