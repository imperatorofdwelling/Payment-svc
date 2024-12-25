package service

import "errors"

var (
	ErrCardAlreadyExists = errors.New("card already exists")
)

// payoutsubscriber errors
var (
	ErrNoNeedToCheck      = errors.New("no need to check")
	ErrCannotStartToCheck = errors.New("cannot start to check")
	ErrEmptyResponse      = errors.New("empty response")
)
