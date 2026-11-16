package service

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestVersionService(t *testing.T) {
	expected := "development_commit_1970-01-01T00:00:00Z"
	assert.Equal(t, expected, GetVersion())
}
