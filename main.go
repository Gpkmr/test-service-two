package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/two", SrvTwo)
	http.HandleFunc("/health", Health)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
		fmt.Println("Defaulting to port", port)
	}

	fmt.Println("Listening on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Panic("Error on startup", err)
	}
}

func SrvTwo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Welcome to service two !\n")
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Service two is running !\n")
	}
}
