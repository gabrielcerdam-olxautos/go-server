package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func HandleHome(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "Hello from home")
}
