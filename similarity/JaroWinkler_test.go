package similarity_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/stringutils/similarity"
)

func TestJaroWinkler(t *testing.T) {
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
