package main

import (
	"fmt"
	"net/http"
)

func requestHandler (w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("This is the new server")
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
