package helpers_test

import (
	"testing"

	"github.com/brazilian-utils/go-brazilian-utils/helpers"
)

func TestGenerateChecksum(t *testing.T) {
	tables := []struct {
		base     string
		weigths  []int
		expected int
	}{
		{"12", []int{10}, 28},
		{"12", []int{10, 9}, 28},
	}

	for _, table := range tables {
		if res := helpers.GenerateChecksum(table.base, table.weigths); res != table.expected {
			t.Errorf("Failed for %s (expected: %d, received: %d)", table.base, table.expected, res)
		}
	}
}
