package root

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"blog-gin_golang_v177/test"
)

func TestIndex(t *testing.T) {
	resp, err := test.HttpCallGet("", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "OK", resp.Message)
	assert.Equal(t, nil, resp.Data)
}
