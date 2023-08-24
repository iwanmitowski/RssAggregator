package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts API Key from headers of HTTP Request
// Example:
// Authorization: ApiKey {key}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authentication info")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("incorrect auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("incorect first part of auth header")
	}

	return vals[1], nil
}
