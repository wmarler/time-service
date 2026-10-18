package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current time is %s\n", time.Now().Format(time.RFC3339))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 6080...")
	http.ListenAndServe(":6080", nil)
}
