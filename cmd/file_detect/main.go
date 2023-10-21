package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := flag.String("dir", ".", "directory to serve")
	keyword := flag.String("keyword", "", "keyword to search")
	flag.Parse()

	if *keyword == "" {
		fmt.Println("keyword is required")
		return
	}

	var files []string
	filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && strings.Contains(info.Name(), *keyword) {
			files = append(files, path)
		}

		return nil
	})

	for _, file := range files {
		fmt.Println("Found file:", file)
	}
}

