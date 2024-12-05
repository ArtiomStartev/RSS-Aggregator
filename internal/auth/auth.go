package auth

import (
	"errors"
	"net/http"
	"strings"
)

const (
	ErrNoAuthenticationInformation = "no authentication information provided"
	ErrMalformedAuthenticationInfo = "malformed authentication information"
)

// GetAPIKey extracts the API key from the headers of an HTTP request
// Example of the header:
// Authorization: ApiKey <API_KEY>
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New(ErrNoAuthenticationInformation)
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New(ErrMalformedAuthenticationInfo)
	}
	if vals[0] != "ApiKey" {
		return "", errors.New(ErrMalformedAuthenticationInfo)
	}

	return vals[1], nil
}
