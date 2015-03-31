package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Usage: go run net-html.go url")
		os.Exit(1)
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Download failed")
		os.Exit(1)
	}
	defer resp.Body.Close()

	n, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Parsing failed")
		os.Exit(1)
	}

	var imgs []string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					imgs = append(imgs, attr.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)

	for _, img := range imgs {
		fmt.Println(img)
	}
}
