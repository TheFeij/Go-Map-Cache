package mapcache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateSize(t *testing.T) {
	value := struct {
		name       string
		age        int
		isMarried  bool
		data       map[string]int
		spouseName *string
	}{
		name:      "abolfazl",
		age:       23,
		isMarried: false,
		data: map[string]int{
			"age": 12,
		},
		spouseName: nil,
	}

	require.Equal(t, int64(8+8+1+0+(3+8)), calculateSize(value))
}
