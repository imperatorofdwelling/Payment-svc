package json

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrType string

var (
	ValidationError    ErrType = "validation_error"
	AuthorizationError ErrType = "authorization_error"
)

type ErrResponse struct {
	Type ErrType `json:"type"`
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
}

type Response struct {
	Data any          `json:"data"`
	Err  *ErrResponse `json:"error"`
}

func Read(r *http.Request, s *any) error {
	return json.NewDecoder(r.Body).Decode(s)
}

func Write(w http.ResponseWriter, status http.ConnState, s any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))

	resp := Response{
		Data: s,
		Err:  nil,
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err.Error())
	}
}

func WriteError(w http.ResponseWriter, status http.ConnState, errMsg string, errType ErrType) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))

	resp := Response{
		Data: nil,
		Err: &ErrResponse{
			Type: errType,
			Code: int(status),
			Msg:  errMsg,
		},
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err.Error())
	}
}