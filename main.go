package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	directoryPtr := flag.String("directory", "pubilc", "directory string to file server")
	flag.Parse()
	fmt.Println("..simple server started at: ", time.Now())
	fmt.Printf("..serving the directory: '%s' on port 8080", *directoryPtr)
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir(*directoryPtr)))
	if err != nil {
		log.Printf("Error running web server for static assets: %v", err)
	}
}
