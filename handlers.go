package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func HandleHome(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "Hello from home")
}

func PostRequest(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}
	response, err := user.toJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "Application/Json")
	w.Write(response)

}
