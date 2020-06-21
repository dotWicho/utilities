package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

// WriteDataToJson saves the content of body in a JSON format file
func WriteDataToJson(body interface{}, fileName string) error {

	if !strings.HasSuffix(fileName, ".json") {
		fileName += ".json"
	}

	file, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return err
	}

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

// LoadDataFromJson reads a JSON format file and returns its contents on body
func LoadDataFromJson(body interface{}, fileName string) error {

	return json.Unmarshal(ReadFile(fileName), body)
}

// WriteDataToYaml saves the content of body in a YAML format file
func WriteDataToYaml(body interface{}, fileName string) error {

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

// LoadDataFromYaml reads a YAML format file and returns its contents on body
func LoadDataFromYaml(body interface{}, fileName string) error {

	return yaml.Unmarshal(ReadFile(fileName), body)
}

// ReadFile reads a file an returns its conte
func ReadFile(fileName string) []byte {

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		panic(errors.New(fmt.Sprintf("%s don't exist, aborting.\n", fileName)))
	} else {
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		return file
	}
	return nil
}
