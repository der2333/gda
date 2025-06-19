package main

import (
	"flag"
	"fmt"
	"os"

	gda "github.com/der2333/gda/src"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: gda <directory>")
		os.Exit(1)
	}

	root := args[0]
	rootSize, detailsInfo, err := gda.GetDirSize(root)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Size of directory '%s': %s\n", root, gda.ReadableSize(rootSize))

	for i := range detailsInfo {
		fmt.Printf("%s: %s\n", detailsInfo[i].Path, gda.ReadableSize(detailsInfo[i].Size))
	}
}
