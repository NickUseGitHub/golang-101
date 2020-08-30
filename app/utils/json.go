package utils

import (
	"encoding/json"
)

// ToJSON makes the response with payload as json format
func ToJSON(payload interface{}) []byte {
	response, err := json.Marshal(payload)
	if err != nil {
		return nil
	}

	return response
}
