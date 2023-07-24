package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		version := os.Getenv("VERSION")
		if version == "" {
			version = "NOT_SET"
		}
		fmt.Fprintf(w, "<h1>Hello world from : %s - App version %s<h1>", host, version)
	})

	http.HandleFunc("/exitWithError", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Exit with Error")
		os.Exit(1)
	})

	http.HandleFunc("/exitSuccess", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Exit Successfully")
		os.Exit(0)
	})

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
