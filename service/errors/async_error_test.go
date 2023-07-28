package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigUserPoolId(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(ErrTodoSkeletonError.Error(), "there are no campaigns assigned to this user")
}

func TestErrorCode(t *testing.T) {
	assert := assert.New(t)

	errorCode := ExtractErrorCode(ErrTodoSkeletonError)

	assert.Equal(*errorCode, 200)
}

func TestFiveHundredErrorCode(t *testing.T) {
	assert := assert.New(t)

	errorCode := ExtractErrorCode(ErrTodoSkeletonError)

	assert.Equal(*errorCode, 400)
}
