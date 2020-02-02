package lib

import (
	"encoding/json"
	"fmt"
	"strings"
)


func URLJoin(base, endpoint string) string {
	//common method for joining url.  I'm sure there is a library for this

	return strings.Join([]string{base, endpoint}, "")
}


func ToJson(resp interface{}) (string, error) {
	//common json encoding method to reduce repetitive code

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("could not create new json object %v", err)
	}

	return string(jsonBytes), nil
}
