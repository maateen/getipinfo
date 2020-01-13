package ipapi

import (
	"encoding/json"
	"net/http"
)

var URL string = "http://ip-api.com/json/"

// GetJson returns GeoLocation of a given IP Address in JSON format.
func GetJson(ipAddress string) map[string]interface{}  {
	resp, err := http.Get(URL + ipAddress)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	var temp interface{}
	json.NewDecoder(resp.Body).Decode(&temp)
	jsonData := temp.(map[string]interface{})

	if jsonData["status"] == "success" {
		return jsonData
	} else {
		panic(jsonData)
	}
}