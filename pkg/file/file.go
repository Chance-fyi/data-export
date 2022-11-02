package file

import (
	"io/ioutil"
	"os"
)

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func Put(data []byte, to string) error {
	return ioutil.WriteFile(to, data, 0644)
}
