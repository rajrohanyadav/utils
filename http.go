package utils

import (
	"encoding/json"
	"net/http"
)

// TODO: Take in an io.Writer to make testing easier
func WriteJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("content-type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
