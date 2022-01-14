package enderecos

import "testing"

type scenario struct {
	input    string
	expected string
}

func TestTipoDeEndereco(t *testing.T) {
	// TestNomeDaFuncao

	scenarios := []scenario{
		{"Rua A", "Rua"},
		{"Avenida A", "Avenida"},
		{"Rodovia A", "Rodovia"},
		{"Batata A", "Não identificado"},
		{"RUA A", "Rua"},
	}

	for _, item := range scenarios {
		result := TipoDeEndereco(item.input)
		if result != item.expected {
			t.Errorf("Expected {%s} but got {%s}", item.expected, result)
		}
	}
}
