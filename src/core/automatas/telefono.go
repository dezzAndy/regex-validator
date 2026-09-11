package automatas

import "regexp"

var regexTelefono = regexp.MustCompile(`^[0-9]{10}$`)

func RegexValidarTelefono(s string) bool {
	return regexTelefono.MatchString(s)
}
