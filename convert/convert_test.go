package convert_test

import (
	"strconv"
	"testing"

	"github.com/akyoto/stringutils/convert"
	qt "github.com/frankban/quicktest"
)

func TestDecToInt(t *testing.T) {
	c := qt.New(t)
	c.Assert(0, qt.Equals, convert.DecToInt([]byte("")))
	c.Assert(0, qt.Equals, convert.DecToInt([]byte("0")))
	c.Assert(1, qt.Equals, convert.DecToInt([]byte("1")))
	c.Assert(10, qt.Equals, convert.DecToInt([]byte("10")))
	c.Assert(100, qt.Equals, convert.DecToInt([]byte("100")))
	c.Assert(123456789, qt.Equals, convert.DecToInt([]byte("123456789")))
	c.Assert(0, qt.Equals, convert.DecToInt([]byte("ZZZ")))
}

func TestHexToInt(t *testing.T) {
	c := qt.New(t)
	c.Assert(0x0, qt.Equals, convert.HexToInt([]byte("")))
	c.Assert(0x0, qt.Equals, convert.HexToInt([]byte("0")))
	c.Assert(0x1, qt.Equals, convert.HexToInt([]byte("1")))
	c.Assert(0xA, qt.Equals, convert.HexToInt([]byte("A")))
	c.Assert(0x10, qt.Equals, convert.HexToInt([]byte("10")))
	c.Assert(0xAFFE, qt.Equals, convert.HexToInt([]byte("AFFE")))
	c.Assert(0xAFFE, qt.Equals, convert.HexToInt([]byte("Affe")))
	c.Assert(0xBADFACE, qt.Equals, convert.HexToInt([]byte("BADFACE")))
	c.Assert(0xBADFACE, qt.Equals, convert.HexToInt([]byte("BadFace")))
	c.Assert(0, qt.Equals, convert.HexToInt([]byte("ZZZ")))
}

func BenchmarkDecToInt(b *testing.B) {
	example := []byte("123456789")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		convert.DecToInt(example)
	}
}

func BenchmarkStrconvDecToInt(b *testing.B) {
	example := "123456789"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseInt(example, 10, 64)
	}
}

func BenchmarkHexToInt(b *testing.B) {
	example := []byte("AFFE")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		convert.HexToInt(example)
	}
}

func BenchmarkStrconvHexToInt(b *testing.B) {
	example := "AFFE"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseInt(example, 16, 64)
	}
}
