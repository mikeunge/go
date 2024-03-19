package filehelper

import (
	"fmt"
	"os"
)

func WriteFile(path string, data string, perm os.FileMode) error {
	return os.WriteFile(path, []byte(data), perm)
}

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read data from file %s, %+v", path, err)
	}

	return data, nil
}
