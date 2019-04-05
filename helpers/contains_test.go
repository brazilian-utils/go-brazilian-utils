package helpers_test

import (
	"testing"

	"github.com/brazilian-utils/go-brazilian-utils/helpers"
)

var sliceOfString []string = []string{"foo", "bar"}

func TestContaints(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	}{
		{"foo", true},
		{"bar", true},
		{"foobar", false},
		{"another string", false},
	}

	for _, table := range tables {
		if res := helpers.Contains(sliceOfString, table.input); res != table.expected {
			t.Errorf("Failed for %s (expected: %t, received: %t)", table.input, table.expected, res)
		}
	}
}
