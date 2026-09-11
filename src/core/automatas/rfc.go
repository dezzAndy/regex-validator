package automatas

import "regexp"

var (
	// Persona física: 4 letras + fecha (YYMMDD) + homoclave (2 alfanuméricos) + dígito verificador
	regexRfcFisica = regexp.MustCompile(`^[A-Z][AEIOUX][A-Z]{2}\d{2}(?:0[1-9]|1[0-2])(?:0[1-9]|[12]\d|3[01])[A-Z\d]{2}\d$`)
	// Persona moral: 3 letras (puede incluir &) + fecha (YYMMDD) + homoclave (3 caracteres)
	regexRfcMoral = regexp.MustCompile(`^[A-ZÑ&]{3}\d{2}(?:0[1-9]|1[0-2])(?:0[1-9]|[12]\d|3[01])[A-Z\d]{3}$`)
)

func RegexValidarRfc(s string) bool {
	return regexRfcFisica.MatchString(s) || regexRfcMoral.MatchString(s)
}
