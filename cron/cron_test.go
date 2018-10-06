package cron

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallAPISuccess(t *testing.T) {
	text, err := Call("0 30 10-13 ? * WED,FRI")
	assert.Equal(t, nil, err)
	assert.Equal(t, "At 30 minutes past the hour, between 10:00 AM and 01:59 PM, only on Wednesday and Friday", text)
}

func TestCallAPIFailed(t *testing.T) {
	text, err := Call("")
	assert.Equal(t, nil, err)
	assert.Equal(t, "Field 'expression' not found.", text)
}
