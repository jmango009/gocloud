package utils

import (
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

var (
	imageTypes = []string{
		"jpeg", "bmp", "gif", "png",
	}
	videoTypes = []string{
		"mpeg", "mp4",
	}
)

type MyCustomClaims struct {
	// This will hold a users username after authenticating.
	// Ignore `json:"username"` it's required by JSON
	Username string `json:"username"`

	// This will hold claims that are recommended having (Expiration, issuer)
	jwt.StandardClaims
}

func IsImage(fileHeader []byte) bool {
	fileType := http.DetectContentType(fileHeader)
	for _, acceptedType := range imageTypes {
		if strings.HasSuffix(fileType, "/"+acceptedType) {
			return true
		}
	}
	return false
}

func IsVideo(fileHeader []byte) bool {
	fileType := http.DetectContentType(fileHeader)
	for _, acceptedType := range videoTypes {
		if strings.HasSuffix(fileType, "/"+acceptedType) {
			return true
		}
	}
	return false
}
