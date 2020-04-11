package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := NewUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJson(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	_ = enc.Encode(data)
}
