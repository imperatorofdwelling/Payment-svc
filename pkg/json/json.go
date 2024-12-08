package json

import (
	"encoding/json"
	"net/http"
)

func Read(r *http.Request, s *interface{}) error {
	return json.NewDecoder(r.Body).Decode(s)
}

func Write(w http.ResponseWriter, status http.ConnState, s any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	return json.NewEncoder(w).Encode(s)
}
