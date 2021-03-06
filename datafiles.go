package utilities

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// WriteDataToJSON saves the content of body in a JSON format file
func WriteDataToJSON(body interface{}, fileName string) error {

	if !strings.HasSuffix(fileName, ".json") {
		fileName += ".json"
	}

	file, _ := json.MarshalIndent(body, "", "  ")

	jsonFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = jsonFile.Write(file)
	if err != nil {
		return err
	}

	return nil
}

// LoadDataFromJSON reads a JSON format file and returns its contents on body
func LoadDataFromJSON(body interface{}, fileName string) error {

	return json.Unmarshal(ReadFile(fileName), body)
}

// WriteDataToYAML saves the content of body in a YAML format file
func WriteDataToYAML(body interface{}, fileName string) error {

	if !strings.HasSuffix(fileName, ".yaml") {
		fileName += ".yaml"
	}

	file, err := yaml.Marshal(body)
	if err != nil {
		return err
	}

	yamlFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = yamlFile.Write(file)
	if err != nil {
		return err
	}

	return nil
}

// LoadDataFromYAML reads a YAML format file and returns its contents on body
func LoadDataFromYAML(body interface{}, fileName string) error {

	return yaml.Unmarshal(ReadFile(fileName), body)
}

// ReadFile reads a file and returns its content if exists, on the other hand nil
func ReadFile(fileName string) []byte {

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		file, err := ioutil.ReadFile(fileName)
		if err == nil {
			return file
		}
	}
	return nil
}

// WriteFile writes io.ReadCloser to a file and returns an error if exists
func WriteFile(body io.Reader, fileName string) error {

	if file, err := os.Create(fileName); err != nil {
		return err
	} else {
		defer func() { _ = file.Close() }()
		if _, err = io.Copy(file, body); err != nil {
			return err
		}
	}

	return nil
}
