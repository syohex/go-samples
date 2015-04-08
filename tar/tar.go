package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tar file")
		os.Exit(1)
	}

	fmt.Println("List: ", os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	tr := tar.NewReader(f)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(hdr.Name)
	}
}
