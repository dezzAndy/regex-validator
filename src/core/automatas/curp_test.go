package automatas

import "testing"

func TestValidarCurp(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado bool
	}{
		{"GARC820714HJCRRL09", true},   // 18 caracteres correctos
		{"GARC820714HJCRRL0", false},   // 17 caracteres (incompleto)
		{"GARC820714HJCRRL091", false}, // 19 caracteres (excedido)
		{"GARC820714HJCRRL0A", false},  // Dígito verificador no numérico
		{"GARC820714HJCRPL09", true},   // Entidad y consonantes válidas (JC=Jalisco)
		{"GARC820714HXXRPL09", false},  // Entidad inexistente (XX)
		{"GHRC820714HJCRRL09", false},  // Segunda letra no vocal
		{"", false},                    // Cadena vacía
	}

	for _, c := range casos {
		resultado := RegexValidarCurp(c.entrada)
		if resultado != c.esperado {
			t.Errorf("Para entrada %q se esperaba %v, pero se obtuvo %v", c.entrada, c.esperado, resultado)
		}
	}
}
