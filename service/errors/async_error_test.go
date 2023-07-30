package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigUserPoolId(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(ErrUnableToStoreFileInS3.Error(), "error storing file in s3")
}

func TestErrorCode(t *testing.T) {
	assert := assert.New(t)

	errorCode := ExtractErrorCode(ErrUnableToStoreFileInS3)

	assert.Equal(*errorCode, 400)
}
