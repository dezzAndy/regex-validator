package automatas

import "testing"

func TestValidarRfc(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado bool
	}{
		{"GARC820714H10", true},  // Persona física completa (13 chars)
		{"GARC820714H1", false},  // Persona física incompleta (12 chars)
		{"GARC820714H1A", false}, // Dígito verificador no numérico
		{"GARC820732H10", false}, // Fecha con día 32
		{"XYZ890101ABC", true},   // Persona moral completa (12 chars)
		{"XYZ890101ABC1", false}, // Persona moral excedida (13 chars)
		{"XYZ890132ABC", false},  // Fecha con día 32 en moral
		{"A&Ñ890101ABC", true},   // Persona moral con caracteres especiales
		{"GARC820714H1a", false}, // Minúsculas
		{"", false},              // Cadena vacía
	}

	for _, c := range casos {
		resultado := RegexValidarRfc(c.entrada)
		if resultado != c.esperado {
			t.Errorf("Para entrada %q se esperaba %v, pero se obtuvo %v", c.entrada, c.esperado, resultado)
		}
	}
}
