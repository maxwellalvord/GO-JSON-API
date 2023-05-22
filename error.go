package main

import "net/http"

var ErrUserInvalid = apiError{Err: "user not valid", Status: http.StatusForbidden}

type apiError struct {
	Err    string
	Status int
}

func (e apiError) Error() string {
	return e.Err
}
