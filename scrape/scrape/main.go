package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	resp, err := http.Get("http://syohex.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get error: %s", err.Error())
		os.Exit(1)
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse error: %s", err.Error())
		os.Exit(1)
	}

	matcher := func(n *html.Node) bool {
		return n.DataAtom == atom.A
	}

	links := scrape.FindAll(node, matcher)
	for i, link := range links {
		fmt.Printf("Link[%2d]: %s\n", i+1, scrape.Attr(link, "href"))
	}
}
