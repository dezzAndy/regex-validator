package automatas

import "regexp"

var regexCorreo = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func RegexValidarCorreo(s string) bool {
	return regexCorreo.MatchString(s)
}
