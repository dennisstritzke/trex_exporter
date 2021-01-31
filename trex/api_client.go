package trex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func queryStatus(address string) (*summary, error) {
	response, err := http.Get(fmt.Sprintf("%s/summary", address))
	if err != nil {
		return nil, err
	}

	var summary summary

	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&summary)
	if err != nil {
		return nil, err
	}

	return &summary, nil
}
