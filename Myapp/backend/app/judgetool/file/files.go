package file

import (
	"os"
	"path/filepath"
)

func saveFile(dir, filename string, data []byte) (string, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	fullPath := filepath.Join(dir, filename)
	err := os.WriteFile(fullPath, data, 0644)
	if err != nil {
		return "", err
	}
	return fullPath, nil
}
