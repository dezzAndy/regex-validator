package main

/*
#include <stdbool.h>
*/
import "C"
import (
	"validador-core/automatas"
)

//export ValidarTelefono
func ValidarTelefono(input *C.char) C.bool {
	return C.bool(automatas.RegexValidarTelefono(C.GoString(input)))
}

//export ValidarCorreo
func ValidarCorreo(input *C.char) C.bool {
	return C.bool(automatas.RegexValidarCorreo(C.GoString(input)))
}

//export ValidarCurp
func ValidarCurp(input *C.char) C.bool {
	return C.bool(automatas.RegexValidarCurp(C.GoString(input)))
}

//export ValidarPass
func ValidarPass(input *C.char) C.bool {
	return C.bool(automatas.ValidarPass(C.GoString(input)))
}

//export ValidarRfc
func ValidarRfc(input *C.char) C.bool {
	return C.bool(automatas.RegexValidarRfc(C.GoString(input)))
}

//export ValidarIP
func ValidarIP(input *C.char) C.bool {
	return C.bool(automatas.RegexValidarIP(C.GoString(input)))
}

//export ValidarCumpleanos
func ValidarCumpleanos(input *C.char) C.bool {
	return C.bool(automatas.ValidarCumpleanos(C.GoString(input)))
}

//export ValidarPlaca
func ValidarPlaca(input *C.char) C.bool {
	return C.bool(automatas.RegexValidarPlaca(C.GoString(input)))
}

func main() {}
