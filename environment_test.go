package utilities

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnvVariable(t *testing.T) {

	t.Run("get default Value if Env don't exists", func(t *testing.T) {

		// Define as var, we used ahead
		key := "TESTING_KEY_GET"
		value := "OneTestString"
		defValue := "default"

		// get Env["TESTING_KEY"]
		result := GetEnvVariable(key, defValue)

		// result must be not equal to value
		assert.NotEqual(t, value, result)

		// result must be equal to defValue
		assert.Equal(t, defValue, result)
	})

	t.Run("get Environment Value if Env exists", func(t *testing.T) {

		// Define as var, we used ahead
		key := "TESTING_KEY_GET"
		value := "OneTestString"
		defValue := "default"

		// Set Env["TESTING_KEY"] = "OneTestString"
		err := os.Setenv(key, value)

		// err must be nil
		assert.Nil(t, err)

		// get Env["TESTING_KEY"]
		result := GetEnvVariable(key, defValue)

		// result must be equal to value
		assert.Equal(t, value, result)
	})
}

func TestSetEnvVariable(t *testing.T) {

	t.Run("get empty value if Env don't exists", func(t *testing.T) {

		// Define as var, we used ahead
		key := "TESTING_KEY_SET"
		value := "default"

		// get Env["TESTING_KEY"]
		result := SetEnvVariable(key, value)

		// result must be not equal to value
		assert.Empty(t, result)
	})

	t.Run("get Environment Value if Env exists", func(t *testing.T) {

		// Define as var, we used ahead
		key := "TESTING_KEY_SET"
		value := "default"
		test := "environment"

		// Set Env["TESTING_KEY"] = "OneTestString"
		err := os.Setenv(key, test)

		// err must be nil
		assert.Nil(t, err)

		// get Env["TESTING_KEY"]
		result := SetEnvVariable(key, value)

		// result must be not equal to value
		assert.NotEmpty(t, result)

		// result must be equals to test's value
		assert.Equal(t, test, result)
	})
}
