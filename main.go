package main

import (
	"fmt"
	"html"
	"net/http"
)

func startAnEchoServer() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.ListenAndServe(":3333", nil)
}

func main() {
	fmt.Println("Start")
	startAnEchoServer()
	fmt.Println("End")
}
