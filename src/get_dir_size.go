package gda

import (
	"os"
	"path/filepath"
)

func GetDirSize(path string) (int64, error) {
	var dirSize int64
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			dirSize += info.Size()
		}
		return nil
	})

	return dirSize, err
}
