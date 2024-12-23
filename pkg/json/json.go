package json

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ErrType string

var (
	ValidationError        ErrType = "validation_error"
	AuthorizationError     ErrType = "authorization_error"
	DecodeBodyError        ErrType = "decode_body_error"
	GettingHeaderDataError ErrType = "getting_header_data_error"
	ExternalApiError       ErrType = "external_api_error"
	InternalApiError       ErrType = "internal_api_error"
	UnmarshallingError     ErrType = "unmarshalling_error"
	ParseError             ErrType = "parse_error"
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

func Read(body io.ReadCloser, s any) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(s)
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
