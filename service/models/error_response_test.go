package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorResponse(t *testing.T) {
	assert := assert.New(t)
	body := `{ "message": "123456789"}`

	errorHttpResponse := ErrorHttpResponse{}
	err := json.Unmarshal([]byte(body), &errorHttpResponse)

	assert.NoError(err)

	assert.Equal(*errorHttpResponse.Message, "123456789")
}
