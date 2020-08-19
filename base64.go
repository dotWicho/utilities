package utilities

import (
	"encoding/base64"
)

// DecodeBase64 returns plain text from base64 entry
func DecodeBase64(encode string) string {

	decoded, _ := base64.StdEncoding.DecodeString(encode)
	return string(decoded)
}

// EncodeBase64 returns base64 encoded text
func EncodeBase64(value string) string {

	return base64.StdEncoding.EncodeToString([]byte(value))
}
