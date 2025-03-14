package main

import (
	"fmt"
	"net/http"
)

func requestHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the new server")
}
