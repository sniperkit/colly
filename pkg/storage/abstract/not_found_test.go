package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundLoginShouldReturnError(t *testing.T) {
	nf := &NotFound{}
	token, err := nf.Login(context.Background())
	assert.NotNil(t, err)
	assert.Equal(t, "storage not found", err.Error())
	assert.Equal(t, "", token)
}
