package common

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func YamlToJson(File string) []byte {
	yamlFile, err := ioutil.ReadFile(File)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Convert YAML to JSON
	var data interface{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Convert map[interface{}]interface{} to map[string]interface{}
	var convertedData interface{}
	switch v := data.(type) {
	case map[interface{}]interface{}:
		convertedData = convertMap(v)
	default:
		convertedData = data
	}

	jsonData, err := json.Marshal(convertedData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//// Write JSON to file
	//err = ioutil.WriteFile(outFile, jsonData, 0644)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}

	return jsonData
}

func convertMap(m map[interface{}]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		key, ok := k.(string)
		if !ok {
			log.Fatalf("error: %v", fmt.Errorf("map key is not a string"))
		}
		switch val := v.(type) {
		case map[interface{}]interface{}:
			result[key] = convertMap(val)
		case []interface{}:
			result[key] = convertSlice(val)
		default:
			result[key] = val
		}
	}
	return result
}

func convertSlice(s []interface{}) []interface{} {
	result := make([]interface{}, len(s))
	for i, v := range s {
		switch val := v.(type) {
		case map[interface{}]interface{}:
			result[i] = convertMap(val)
		case []interface{}:
			result[i] = convertSlice(val)
		default:
			result[i] = val
		}
	}
	return result
}
