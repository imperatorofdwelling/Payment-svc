package v1

import "errors"

var (
	ErrGettingIdempotenceKey = errors.New("error getting idempotence key")
	ErrUnmarshallingBody     = errors.New(`error unmarshalling body`)
)
