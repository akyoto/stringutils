package convert_test

import (
	"strconv"
	"testing"

	"github.com/akyoto/stringutils/convert"
	"github.com/stretchr/testify/assert"
)

func TestDecToInt(t *testing.T) {
	assert.Equal(t, 0, convert.DecToInt([]byte("")))
	assert.Equal(t, 0, convert.DecToInt([]byte("0")))
	assert.Equal(t, 1, convert.DecToInt([]byte("1")))
	assert.Equal(t, 10, convert.DecToInt([]byte("10")))
	assert.Equal(t, 100, convert.DecToInt([]byte("100")))
	assert.Equal(t, 123456789, convert.DecToInt([]byte("123456789")))
	assert.Equal(t, 0, convert.DecToInt([]byte("ZZZ")))
}

func TestHexToInt(t *testing.T) {
	assert.Equal(t, 0x0, convert.HexToInt([]byte("")))
	assert.Equal(t, 0x0, convert.HexToInt([]byte("0")))
	assert.Equal(t, 0x1, convert.HexToInt([]byte("1")))
	assert.Equal(t, 0xA, convert.HexToInt([]byte("A")))
	assert.Equal(t, 0x10, convert.HexToInt([]byte("10")))
	assert.Equal(t, 0xAFFE, convert.HexToInt([]byte("AFFE")))
	assert.Equal(t, 0xAFFE, convert.HexToInt([]byte("Affe")))
	assert.Equal(t, 0xBADFACE, convert.HexToInt([]byte("BADFACE")))
	assert.Equal(t, 0xBADFACE, convert.HexToInt([]byte("BadFace")))
	assert.Equal(t, 0, convert.HexToInt([]byte("ZZZ")))

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
