package gda

import "fmt"

func buildResult(root string, rootSize int64, detailsInfo []DirInfo) {
	fmt.Printf("Size of directory '%s': %s\n", root, readableSize(rootSize))

	for i := range detailsInfo {
		fmt.Printf("%s: %s\n", detailsInfo[i].Path, readableSize(detailsInfo[i].Size))
	}
}
