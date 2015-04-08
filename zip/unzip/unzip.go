package main

import (
	"archive/zip"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: unzip file")
		os.Exit(1)
	}

	r, err := zip.OpenReader(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	defer r.Close()

	for _, zf := range r.File {
		fmt.Println(zf.Name)
	}
}
