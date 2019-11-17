package similarity_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/stringutils/similarity"
)

func TestJaroWinkler(t *testing.T) {
	tests := []struct {
		A          string
		B          string
		Similarity float64
	}{
		{"AL", "AL", 1.0},
		{"AAAAAAAAABCCCC", "AAAAAAAAABCCCC", 1},
		{"ABCVWXYZ", "CABVWXYZ", 0.9625},
		{"MARTHA", "MARHTA", 0.9611111111111111},
		{"JONES", "JOHNSON", 0.8323809523809523},
		{"A", "B", 0},
		{"ABCDEF", "123456", 0},
	}

	for _, test := range tests {
		s := similarity.JaroWinkler(test.A, test.B)
		assert.Equal(t, s, test.Similarity)
	}
}

func TestDescendingSimilarity(t *testing.T) {
	descendingSimilarity := []string{
		"Hello World",
		"Hello Worl",
		"Hello Wor",
		"Hello Wo",
		"Hello W",
		"Hello ",
		"Hello",
		"Hell",
		"Hel",
		"He",
		"H",
	}

	previousValue := 1.1

	for _, test := range descendingSimilarity {
		assert.Equal(t, similarity.JaroWinkler(test, test), 1.0)
		value := similarity.JaroWinkler(descendingSimilarity[0], test)

		if value >= previousValue {
			t.Fatalf("Similarity %f must be lower than the previous value %f", value, previousValue)
		}

		previousValue = value
	}
}
