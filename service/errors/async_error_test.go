package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigUserPoolId(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(ErrTodoSkeletonError.Error(), "todo fill in the error message")
}

func TestErrorCode(t *testing.T) {
	assert := assert.New(t)

	errorCode := ExtractErrorCode(ErrTodoSkeletonError)

	assert.Equal(*errorCode, 400)
}
