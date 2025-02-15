package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var startAt = time.Now()

func main() {
	http.HandleFunc("/helthz", Helthz)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, %s, You are %s years old.", name, age)
}

func Helthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startAt)
	if duration.Seconds() < 10 || duration.Seconds() > 30 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Error: Server is not healthy")))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("OK: Server is healthy"))
	}
}

