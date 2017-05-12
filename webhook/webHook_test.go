package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHook(t *testing.T) {
	hook := NewHook("123", []byte("nill"))
	assert.Equal(t, hook.secret, []byte("nill"))
	assert.Equal(t, hook.verifyToken, "123")
}
