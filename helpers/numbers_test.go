package helpers_test

import (
	"testing"

	"github.com/brazilian-utils/go-brazilian-utils/helpers"
)

func TestOnlyNumbers(t *testing.T) {
	tables := []struct {
		input  string
		output string
	}{
		{"test", ""},
		{"123test", "123"},
		{"test123", "123"},
		{"123test123", "123123"},
	}

	for _, table := range tables {
		formated := helpers.OnlyNumbers(table.input)

		if formated != table.output {
			t.Errorf("Ouput invalid, given %v, expected %v", formated, table.output)
		}
	}
}
