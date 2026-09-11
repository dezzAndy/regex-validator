package automatas

import "testing"

func TestValidarPlaca(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado bool
	}{
		{"ABC-1234", true},
		{"XYZ-0000", true},
		{"ABC-123", false},   // Solo 3 dígitos
		{"ABCD-1234", false}, // 4 letras
		{"A12-3456", false},  // Letras y números mezclados
		{"ABC-12345", false}, // 5 dígitos
		{"abc-1234", false},  // Minúsculas
		{"ABC1234", false},   // Sin guion
		{"", false},
	}

	for _, c := range casos {
		resultado := RegexValidarPlaca(c.entrada)
		if resultado != c.esperado {
			t.Errorf("Para entrada %q se esperaba %v, pero se obtuvo %v", c.entrada, c.esperado, resultado)
		}
	}
}
