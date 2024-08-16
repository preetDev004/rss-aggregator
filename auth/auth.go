package auth

import (
	"errors"
	"net/http"
	"strings"
)

// getting the apikey from the header of http request and getting the user from database
// example:
// Authorization: ApiKey {insert apikey here}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" || len(vals[1]) != 64 {
		return "", errors.New("malformed auth header")
	}
	return vals[1], nil
}
