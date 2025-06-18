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
		fmt.Printf("Usage: gda <directory>")
		os.Exit(1)
	}

	targetDir := args[0]
	dirSize, err := gda.GetDirSize(targetDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Size of directory '%s': %s\n", targetDir, gda.ReadableSize(dirSize))
}
