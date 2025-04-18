package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 2
	// )

	// actual := AddOne(input)
	// if actual != output {
	// 	t.Errorf("AddOne(%d) = %d, want %d", 1, actual, output)
	// }

	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
}

func TestAddOneRequire(t *testing.T) {
	// block next text if require failed
	require.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
}

// go test
// go test --coverprofile=coverage.out
// go tool cover -html=coverage.out -o coverage.html
