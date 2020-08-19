package utilities

import (
	"os"
)

// GetEnvVariable try to get an environment variable, returns default value otherwise
func GetEnvVariable(name, defVal string) string {

	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	return defVal
}

// SetEnvVariable try to set an environment variable, returns current value if exist, empty string otherwise
func SetEnvVariable(name, value string) string {

	val, ok := os.LookupEnv(name)
	_ = os.Setenv(name, value)

	if ok {
		return val
	}
	return ""
}