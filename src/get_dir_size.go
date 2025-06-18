package gda

import (
	"io/fs"

	"github.com/charlievieth/fastwalk"
)

func GetDirSize(path string) (int64, error) {
	var dirSize int64

	conf := fastwalk.Config{}

	walkFn := func(path string, target fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !target.IsDir() {
			info, err := target.Info()
			if err != nil {
				return err
			}
			dirSize += info.Size()
		}
		return nil
	}
	err := fastwalk.Walk(&conf, path, walkFn)

	return dirSize, err
}
