package automatas

import (
	"strings"
	"unicode"
)

// var RegexPass = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()-_=+[]{}\|;:,<.>/?]`)
// var RegexPass = regexp.MustCompile(`^(?=.*a-z)(?=.*A-Z)(?=.*0-9)(?=.*!-/)(?=.*:-@)[a-zA-Z0-9!-/:-@]{8,}`)
// var RegexPass = regexp.MustCompile(`^[a-zA-Z0-9!-/:-@]{8,}`)

func ValidarPass(pass string) bool {
	if len(pass) < 8 {
		return false
	}

	var (
		tieneMayuscula bool
		tieneNumero    bool
		tieneEspecial  bool
	)

	// Caracteres especiales admitidos
	caracteresEspeciales := "!@#$%^&*()-_=+[]{}|;:,.<>?/~`"

	for _, ch := range pass {
		switch {
		case unicode.IsUpper(ch):
			tieneMayuscula = true
		case unicode.IsDigit(ch):
			tieneNumero = true
		case strings.ContainsRune(caracteresEspeciales, ch):
			tieneEspecial = true
		}

		// Si ya se cumplen todas las condiciones, podemos terminar antes si no hay más restricciones
		if tieneMayuscula && tieneNumero && tieneEspecial {
			return true
		}
	}

	return tieneMayuscula && tieneNumero && tieneEspecial
}
