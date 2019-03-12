package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root, _ := filepath.Abs(".")
	fmt.Println("Processing path", root)

	filepath.Walk(root, walkPath)
}

func walkPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory", path)
		} else {
			fmt.Println("File", path)
		}
	}

	return nil
}