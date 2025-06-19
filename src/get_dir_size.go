package gda

import (
	"io/fs"
	"path/filepath"
	"sync"

	"github.com/charlievieth/fastwalk"
)

type DirInfo struct {
	Path string
	Size int64
}

func GetDirSize(root string) (int64, []DirInfo, error) {
	// 对root目录进行规范化处理
	root = filepath.Clean(root)

	var rootSize int64

	detailsInfo := make([]DirInfo, 0) // 存储root的子目录信息
	detailsSize := make(map[string]int64)

	var mu sync.Mutex

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

			mu.Lock()
			rootSize += info.Size()

			// 将文件信息添加到detailsInfo中
			dir := filepath.Dir(path)
			for dir != root {
				// 获取到root的子目录
				if filepath.Dir(dir) != root {
					dir = filepath.Dir(dir)
				} else {

					detailsSize[dir] += info.Size()

					break
				}
			}
			mu.Unlock()

		} else if filepath.Dir(path) == root {
			mu.Lock()
			detailsInfo = append(detailsInfo, DirInfo{Path: path, Size: 0})
			detailsSize[path] = 0
			mu.Unlock()
		}

		return nil
	}

	err := fastwalk.Walk(&conf, root, walkFn)

	for i := range detailsInfo {
		size := detailsSize[detailsInfo[i].Path]
		detailsInfo[i].Size = size
	}

	return rootSize, detailsInfo, err
}
