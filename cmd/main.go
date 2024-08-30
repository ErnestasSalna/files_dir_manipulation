package main

import (
	"embed"
	s "files_dir_scanner/scanner"
	"fmt"
	"os"
)

var exampleFS embed.FS

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory path.")
		return
	}

	path := os.Args[1] // Get path from the command line
	entries, err := s.ScanPath(exampleFS, path)
	if err != nil {
		fmt.Printf("Error scanning path %s: %v\n", path, err)
		return
	}

	fmt.Printf("Contents of %s:\n", path)
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}
