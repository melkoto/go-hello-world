package main

import (
	"fmt"
	"hello-world/pkg/handlers"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting port on %s", port))

	_ = http.ListenAndServe(port, nil)
}
