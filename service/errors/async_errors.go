package errors

import "errors"

var ErrUnableToStoreFileInS3 = errors.New("error storing file in s3")
var ErrMarshalingResponseFromDB = errors.New("error marshaling response from db")

func ExtractErrorCode(err error) *int {
	var errorCode int

	switch {
	case errors.Is(err, ErrUnableToStoreFileInS3):
		errorCode = 400
	default:
		errorCode = 500
	}

	return &errorCode
}
