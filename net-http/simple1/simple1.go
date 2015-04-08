package main

import (
	"fmt"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/hello", helloHandle)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return
	}

	return
}
