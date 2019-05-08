package unsafe_test

import (
	"testing"

	"github.com/akyoto/stringutils/unsafe"
	"github.com/stretchr/testify/assert"
)

func TestBytesToString(t *testing.T) {
	out := "hello"
	in := []byte(out)

	assert.Equal(t, out, unsafe.BytesToString(in))
}

func TestStringToBytes(t *testing.T) {
	in := "hello"
	out := []byte(in)

	assert.Equal(t, out, unsafe.StringToBytes(in))
}
