package errors

import "errors"

// TODO
var ErrTodoSkeletonError = errors.New("todo fill in the error message")

func ExtractErrorCode(err error) *int {
	var errorCode int

	switch {
	case errors.Is(err, ErrTodoSkeletonError):
		errorCode = 400
	default:
		errorCode = 500
	}

	return &errorCode
}
