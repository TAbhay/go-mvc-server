package utils

import (
	"os"
	"path/filepath"
)

func SaveToFile(filename string, data []byte) error {

	filename = filepath.Join("reports", filename)
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
