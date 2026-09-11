package automatas

import "testing"

func TestValidarTelefono(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado bool
	}{
		{"3312345678", true},   // 10 dígitos válidos
		{"123456789", false},   // 9 dígitos (incompleto)
		{"12345678901", false}, // 11 dígitos (excedido)
		{"331234567a", false},  // Letras intermedias
		{"", false},            // Cadena vacía
		{"33-1234567", false},  // Símbolos no numéricos
	}

	for _, c := range casos {
		resultado := RegexValidarTelefono(c.entrada)
		if resultado != c.esperado {
			t.Errorf("Para entrada %q se esperaba %v, pero se obtuvo %v", c.entrada, c.esperado, resultado)
		}
	}
}
