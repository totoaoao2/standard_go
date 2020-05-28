package hello

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Hello(t *testing.T) {
	t.Run("hello", func(t *testing.T) {
		actual := Output()
		assert.Equal(t, actual, "Hello world")
	})
}