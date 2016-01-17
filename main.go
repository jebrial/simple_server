package main

import (
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("No folder specified")
	}
	folder := os.Args[1]
	http.Handle("/", http.FileServer(http.Dir(folder)))
	http.ListenAndServe(":8080", nil)
}
