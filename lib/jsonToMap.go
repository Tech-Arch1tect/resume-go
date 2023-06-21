package lib

import (
	"encoding/json"
	"os"
)

func GetMapFromJson(jsonFile string) map[string]interface{} {
	// get json string from file
	jsonStr, err := os.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	// unmarshal json into map
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}
	return jsonMap
}
