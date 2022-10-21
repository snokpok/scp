# SCP Backend written in Golang

Using golang because I like it; also kind of an exercise

- Server's required because we need a way to know what the current user's valid access token is; it's not possible to do this on the frontend because browser clients are stateless and would require the user to authenticate every time
- Basically this provides an endpoint for the client to connect to a specific user's SCP, and the refreshing of token is done on the backend & updated in the database every time

Database: MongoDB;

A few endpoints:

- POST /user: create user with {email, username, access_token, refresh_token}
- GET /me: get the user's email username, current access token and refresh token
- GET /scp: get currently playing song data (from spotify api)
  - (backend): will refresh token to get new access_token & update in db upon request if current access_token expired upon request to Spotify API
  - (this is applied to all routes so something like a middleware that is)
- GET /login: login user, get back access token to access current api to get SCP

## Installation

1. Clone this repo to local
2. Run `go mod download` to download dependencies to local cache
3. Retrieve private files: `.env` and the MongoDB user access key file to be placed in `src/certs/` as `cert-rw-user.pem`
4. Run the server in `release` mode by `go run main.go`. To build the server into binary run `go build .` instead. To change the server running mode change the `DEPLOY_ENV` variable in your `.env` file. The other acceptable value is `debug` which runs Gin in debug mode.
