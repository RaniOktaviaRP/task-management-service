package helper

import (
	"encoding/json"
	"net/http"
)

// ReadFromRequestBody membaca body request dan decode ke struct
func ReadFromRequestBody(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

// WriteToResponseBody menulis response ke w
func WriteToResponseBody(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
