package automatas

import "testing"

func TestValidarCorreo(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado bool
	}{
		{"hola@gmail.com", true},         // 10 dígitos válidos
		{"1puto", false},                 // 9 dígitos (incompleto)
		{"@gmail.com", false},            // 11 dígitos (excedido)
		{"holajose@mamas.h", false},      // Letras intermedias
		{"", false},                      // Cadena vacía
		{"holaaa-_1@hotmail.coom", true}, // Símbolos no numéricos
	}

	for _, c := range casos {
		resultado := RegexValidarCorreo(c.entrada)
		if resultado != c.esperado {
			t.Errorf("Para entrada %q se esperaba %v, pero se obtuvo %v", c.entrada, c.esperado, resultado)
		}
	}
}
