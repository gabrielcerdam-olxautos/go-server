package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) toJson() ([]byte, error) {
	return json.Marshal(u)
}

type MetaData interface{}
