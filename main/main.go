package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		returnString := "Your request was received on: "
		currTime := time.Now()
		w.Write([]byte(returnString))
		w.Write([]byte(currTime.Format(time.RFC3339)))
		fmt.Println("Received request in golang server.")
	})
	fmt.Println("Listening to 8080...")
	http.ListenAndServe(":8080", nil)
}
