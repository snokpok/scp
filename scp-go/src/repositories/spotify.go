package repositories

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func RequestSCPFromSpotify(accessToken string) (map[string]interface{}, error) {
	var resultScp map[string]interface{}
	scpUrl := "https://api.spotify.com/v1/me/player/currently-playing"
	hcli := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, scpUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := hcli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseBody, &resultScp)
	if err != nil {
		return nil, err
	}
	return resultScp, nil
}

func RequestNewAccessTokenFromSpotify(refreshToken string) (string, error) {
	form := url.Values{}
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", refreshToken)

	refreshUrl := "https://accounts.spotify.com/api/token"
	reqRefreshToken, err := http.NewRequest(http.MethodPost, refreshUrl, strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	encodedHeaderClient := base64.StdEncoding.EncodeToString([]byte(os.Getenv("SPOTIFY_CLIENT_ID") + ":" + os.Getenv("SPOTIFY_CLIENT_SECRET")))
	reqRefreshToken.Header.Set("Authorization", "Basic "+encodedHeaderClient)
	reqRefreshToken.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	hcli := &http.Client{}
	resp, err := hcli.Do(reqRefreshToken)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println(string(responseBody))
	resultNewSpotifyToken := make(map[string]interface{})
	err = json.Unmarshal(responseBody, &resultNewSpotifyToken)
	if err != nil {
		return "", err
	}
	newAcTkn := resultNewSpotifyToken["access_token"].(string)
	log.Println(newAcTkn)
	return newAcTkn, nil
}
