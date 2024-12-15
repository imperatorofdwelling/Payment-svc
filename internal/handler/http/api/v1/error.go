package v1

import "errors"

var (
	ErrGettingIdempotenceKey        = errors.New("error getting idempotence key")
	ErrUnmarshallingPaymentResponse = errors.New("error unmarshalling payment response")
)
