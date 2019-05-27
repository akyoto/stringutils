package unsafe_test

import (
	"testing"

	"github.com/akyoto/stringutils/unsafe"
	qt "github.com/frankban/quicktest"
)

func TestBytesToString(t *testing.T) {
	out := "hello"
	in := []byte(out)
	c := qt.New(t)
	c.Assert(out, qt.DeepEquals, unsafe.BytesToString(in))
}

func TestStringToBytes(t *testing.T) {
	in := "hello"
	out := []byte(in)
	c := qt.New(t)
	c.Assert(out, qt.DeepEquals, unsafe.StringToBytes(in))
}
