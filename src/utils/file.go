package utils

import (
	"errors"
	"os"
)

var GlobalPath string

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func FilePath(path string) string {
	return GlobalPath + path
}
