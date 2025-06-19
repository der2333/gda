package gda

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/charlievieth/fastwalk"
)

type DirInfo struct {
	Path string
	Size int64
}

func GetDirSize(root string) (int64, error) {
	// 对root目录进行规范化处理
	root = filepath.Clean(root)

	var rootSize int64

	// 存储root的子目录信息
	detailsInfo := make([]DirInfo, 0)
	detailsSize := make(map[string]int64)

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
			rootSize += info.Size()

			// 将文件信息添加到details中
			dir := filepath.Dir(path)
			for dir != root {
				if filepath.Dir(dir) != root {
					dir = filepath.Dir(dir)
					// fmt.Println(dir)
				} else {
					_, exist := detailsSize[dir]
					if !exist {
						// 如果details中没有该目录，则添加
						detailsSize[dir] = info.Size()
						detailsInfo = append(detailsInfo, DirInfo{Path: dir, Size: 0})
					} else {
						// 如果details中有该目录，则更新大小
						detailsSize[dir] += info.Size()
						for i, detail := range detailsInfo {
							if detail.Path == dir {
								detailsInfo[i].Size = detailsSize[dir]
								break
							}
						}
					}
					break
				}
				// 更新details中对应目录的大小
				// for _, detail := range details {
				// 	if detail.Path == dir {
				// 		detail.Size += info.Size()
				// 		break
				// 	}
				// }
			}
		}
		// else if filepath.Dir(path) == root {
		// 	// 初始化root的所有子目录
		// 	details = append(details, DirInfo{Path: path, Size: 0})
		// }

		return nil
	}

	err := fastwalk.Walk(&conf, root, walkFn)

	fmt.Println(detailsInfo)

	return rootSize, err
}
