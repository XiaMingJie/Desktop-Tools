package utils

import "os"

func IsFileExists(path string) bool {
	_,err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}