package validators

import (
	"errors"
	"strconv"
)

func ValidateMessageMap(data map[string][]string) error {
	// Check that the map has the correct keys
	if _, ok := data["content"]; !ok {
		return errors.New("no content")
	}
	if _, ok := data["unlocks_at"]; !ok {
		return errors.New("no unlocks_at")
	}
	// Check that it has no more keys than necessary
	for k := range data {
		if k != "content" && k != "unlocks_at" {
			return errors.New("invalid key: " + k)
		} else if k == "unlocks_at" {
			if len(data[k]) != 1 {
				return errors.New("no value or multiple values for unlocks_at")
			}
			_, err := strconv.Atoi(data[k][0])
			if err != nil {
				return errors.New("invalid value for key: " + k)
			}
		} else if k == "content" {
			if len(data[k]) != 1 {
				return errors.New("invalid value for key: " + k)
			}
		}
	}
	return nil
}