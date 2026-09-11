package automatas

import "testing"

func TestValidarCumpleanos(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado bool
	}{
		{"15/08/1990", true},  // Fecha típica
		{"29/02/2000", true},  // Bisiesto divisible entre 400
		{"29/02/2024", true},  // Bisiesto divisible entre 4
		{"29/02/1900", false}, // No bisiesto (divisible entre 100 pero no 400)
		{"29/02/2023", false}, // Año no bisiesto
		{"31/04/2020", false}, // Abril tiene 30 días
		{"31/12/2025", true},  // Mes con 31 días
		{"00/01/2020", false}, // Día cero
		{"10/13/2020", false}, // Mes 13
		{"10/00/2020", false}, // Mes cero
		{"1/01/2020", false},  // Día de un dígito
		{"15/08/90", false},   // Año de dos dígitos
		{"15-08-1990", false}, // Separador incorrecto
		{"", false},
	}

	for _, c := range casos {
		resultado := ValidarCumpleanos(c.entrada)
		if resultado != c.esperado {
			t.Errorf("Para entrada %q se esperaba %v, pero se obtuvo %v", c.entrada, c.esperado, resultado)
		}
	}
}
