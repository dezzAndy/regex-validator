package automatas

import "regexp"

var regexPlaca = regexp.MustCompile(`^[A-Z]{3}-\d{4}$`)

func RegexValidarPlaca(s string) bool {
	return regexPlaca.MatchString(s)
}
