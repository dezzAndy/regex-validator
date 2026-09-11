package automatas

import "testing"

func TestValidarPass(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado bool
	}{
		{"9aZsp7WxjvNkFdhfrgtvTs6XS4vXHc*eB#G4$5SJZR8fjgR0^", true}, // Mayúscula, número y símbolo
		{"mHYzY2071efD5nXJUem#", true},                              // Mayúscula, número y símbolo
		{"u*1s!tYPgvPVa8XtD6WmCcs@Rn*", true},                       // Mayúscula, número y símbolo
		{"SR8j3w2nzm0j6@#bWcU#g6xgZ@W!88XjKn1", true},               // Mayúscula, número y símbolo
		{"holahola", false},                                         // Sin mayúscula, número ni símbolo
		{"33-1234567", false},                                       // Sin mayúscula
		{"HOLA1234", false},                                         // Sin símbolo
		{"Hola123!", true},                                          // Condiciones mínimas (8 chars)
		{"Hola12", false},                                           // Menos de 8 caracteres
	}

	for _, c := range casos {
		resultado := ValidarPass(c.entrada)
		if resultado != c.esperado {
			t.Errorf("Para entrada %q se esperaba %v, pero se obtuvo %v", c.entrada, c.esperado, resultado)
		}
	}
}
