package utils

import (
	"os"
)

var GlobalPath string

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
