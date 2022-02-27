package helper

import (
	"errors"
	"strings"
)

func CheckBookStatus(status string) (bool, error) {
	var stat bool
	switch strings.ToLower(status) {
	case "not available":
		stat = false
	case "available":
		stat = true
	default:
		return stat, errors.New(`book status can either be 'available' or 'not available'`)
	}
	return stat, nil
}
