package utils

import (
	"encoding/json"
	"net/http"
)

type H map[string]interface{}

func EncodeJosn(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
