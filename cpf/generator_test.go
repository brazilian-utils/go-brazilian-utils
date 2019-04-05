package cpf_test

import (
	"testing"

	"github.com/brazilian-utils/go-brazilian-utils/cpf"
)

func TestGenerator(t *testing.T) {
	generated := cpf.Generate()

	if len(generated) != 11 {
		t.Errorf("the generated CPF does not have the correct length of 11: received %d", len(generated))
	}

	if !cpf.IsValid(generated) {
		t.Errorf("the generated CPF %s is not valid", generated)
	}
}

func BenchmarkGeneratorWithoutMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generated := cpf.Generate()

		b.StopTimer()

		if !cpf.IsValid(generated) {
			b.Errorf("the generated CPF %s is not valid", generated)
		}

		b.StartTimer()
	}
}

// TODO: add generation based on state
