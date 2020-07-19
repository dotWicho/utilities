package utilities

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

type UserInfo struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func TestWriteDataToJSON(t *testing.T) {

	t.Run("Must be equals file and expected if fileName dont have extension", func(t *testing.T) {
		// Define our data to work
		fileName := "example"
		User := &UserInfo{Name: "Jonah Doe", Age: 45}
		expected := []byte{0x7b, 0xa, 0x20, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4a, 0x6f, 0x6e, 0x61, 0x68, 0x20, 0x44, 0x6f, 0x65, 0x22, 0x2c, 0xa, 0x20, 0x20, 0x22, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x35, 0xa, 0x7d}

		// fire up
		writeErr := WriteDataToJSON(User, fileName)
		defer os.Remove(fileName + ".json")

		// if ok?
		assert.Nil(t, writeErr)

		// Read content of file
		file, err := ioutil.ReadFile(fileName + ".json")

		// not errors?
		assert.Nil(t, err)

		// Was write correctly?
		assert.Equal(t, expected, file)
	})

	t.Run("Must be equals file and expected if fileName have extension", func(t *testing.T) {
		// Define our data to work
		fileName := "example.json"
		User := &UserInfo{Name: "Jonah Doe", Age: 45}
		expected := []byte{0x7b, 0xa, 0x20, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4a, 0x6f, 0x6e, 0x61, 0x68, 0x20, 0x44, 0x6f, 0x65, 0x22, 0x2c, 0xa, 0x20, 0x20, 0x22, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x35, 0xa, 0x7d}

		// fire up
		writeErr := WriteDataToJSON(User, fileName)
		defer os.Remove(fileName)

		// if ok?
		assert.Nil(t, writeErr)

		// Read content of file
		file, err := ioutil.ReadFile(fileName)

		// not errors?
		assert.Nil(t, err)

		// Was write correctly?
		assert.Equal(t, expected, file)
	})
}

func TestWriteDataToYAML(t *testing.T) {

	t.Run("Must be equals file and expected if fileName dont have extension", func(t *testing.T) {
		// Define our data to work
		fileName := "example"
		User := &UserInfo{Name: "Jonah Doe", Age: 45}
		expected := []byte{0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0x4a, 0x6f, 0x6e, 0x61, 0x68, 0x20, 0x44, 0x6f, 0x65, 0xa, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x34, 0x35, 0xa}

		// fire ups
		writeErr := WriteDataToYAML(User, fileName)
		defer os.Remove(fileName + ".yaml")

		// if ok?
		assert.Nil(t, writeErr)

		// Read content of file
		file, err := ioutil.ReadFile(fileName + ".yaml")

		// not errors?
		assert.Nil(t, err)

		// Was write correctly?
		assert.Equal(t, expected, file)
	})

	t.Run("Must be equals file and expected if fileName have extension", func(t *testing.T) {
		// Define our data to work
		fileName := "example.yaml"
		User := &UserInfo{Name: "Jonah Doe", Age: 45}
		expected := []byte{0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0x4a, 0x6f, 0x6e, 0x61, 0x68, 0x20, 0x44, 0x6f, 0x65, 0xa, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x34, 0x35, 0xa}

		// fire ups
		writeErr := WriteDataToYAML(User, fileName)
		defer os.Remove(fileName)

		// if ok?
		assert.Nil(t, writeErr)

		// Read content of file
		file, err := ioutil.ReadFile(fileName)

		// not errors?
		assert.Nil(t, err)

		// Was write correctly?
		assert.Equal(t, expected, file)
	})

}

func TestLoadDataFromJSON(t *testing.T) {

	// Define our data to work
	fileName := "example.json"
	User := []byte{0x7b, 0xa, 0x20, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4a, 0x6f, 0x6e, 0x61, 0x68, 0x20, 0x44, 0x6f, 0x65, 0x22, 0x2c, 0xa, 0x20, 0x20, 0x22, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x35, 0xa, 0x7d}
	UserReaded := &UserInfo{}

	// We create out file to read as JSON
	err := ioutil.WriteFile(fileName, User, 0644)

	// if not errors?
	assert.Nil(t, err)

	// fire up
	readError := LoadDataFromJSON(UserReaded, fileName)
	defer os.Remove(fileName)

	//
	assert.Nil(t, readError)

	//
	assert.Equal(t, "Jonah Doe", UserReaded.Name)
	assert.Equal(t, 45, UserReaded.Age)
}

func TestLoadDataFromYAML(t *testing.T) {

	// Define our data to work
	fileName := "example.yaml"
	User := []byte{0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0x4a, 0x6f, 0x6e, 0x61, 0x68, 0x20, 0x44, 0x6f, 0x65, 0xa, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x34, 0x35, 0xa}
	UserReaded := &UserInfo{}

	// We create out file to read as YAML
	err := ioutil.WriteFile(fileName, User, 0644)

	// if not errors?
	assert.Nil(t, err)

	// fire up
	readError := LoadDataFromYAML(UserReaded, fileName)
	defer os.Remove(fileName)

	//
	assert.Nil(t, readError)

	//
	assert.Equal(t, "Jonah Doe", UserReaded.Name)
	assert.Equal(t, 45, UserReaded.Age)
}

func TestReadFile(t *testing.T) {

	// Define our data to work
	fileName1 := "example.yaml"
	User1 := []byte{0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0x4a, 0x6f, 0x6e, 0x61, 0x68, 0x20, 0x44, 0x6f, 0x65, 0xa, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x34, 0x35, 0xa}

	// We create out file to read as YAML
	err1 := ioutil.WriteFile(fileName1, User1, 0644)

	// if not errors?
	assert.Nil(t, err1)

	// fire up
	result1 := ReadFile(fileName1)
	defer os.Remove(fileName1)

	//
	assert.Equal(t, User1, result1)

	// Again, define our data to work
	fileName2 := "example.json"
	User2 := []byte{0x7b, 0xa, 0x20, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4a, 0x6f, 0x6e, 0x61, 0x68, 0x20, 0x44, 0x6f, 0x65, 0x22, 0x2c, 0xa, 0x20, 0x20, 0x22, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x35, 0xa, 0x7d}

	// We create out file to read as JSON
	err2 := ioutil.WriteFile(fileName2, User2, 0644)

	// if not errors?
	assert.Nil(t, err2)

	// fire up
	result2 := ReadFile(fileName2)
	defer os.Remove(fileName2)

	//
	assert.Equal(t, User2, result2)
}
