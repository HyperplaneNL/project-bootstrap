package schema

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
)

// FromYamlFile - create a schema from a yaml file.
func FromYamlFile(file string) (*Schema, error) {
	yamlString, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var schema Schema
	err = yaml.Unmarshal(yamlString, &schema)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}

// FromJsonFile - create a schema from a json file.
func FromJsonFile(file string) (*Schema, error) {
	jsonString, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var schema Schema
	err = json.Unmarshal(jsonString, &schema)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
