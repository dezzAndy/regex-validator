package automatas

import (
	"regexp"
	"strconv"
)

var regexCumpleanos = regexp.MustCompile(`^(\d{2})/(\d{2})/(\d{4})$`)

func ValidarCumpleanos(s string) bool {
	m := regexCumpleanos.FindStringSubmatch(s)
	if m == nil {
		return false
	}

	dia, _ := strconv.Atoi(m[1])
	mes, _ := strconv.Atoi(m[2])
	anio, _ := strconv.Atoi(m[3])

	if mes < 1 || mes > 12 {
		return false
	}

	diasPorMes := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if esBisiesto(anio) {
		diasPorMes[1] = 29
	}

	return dia >= 1 && dia <= diasPorMes[mes-1]
}

func esBisiesto(anio int) bool {
	return anio%4 == 0 && (anio%100 != 0 || anio%400 == 0)
}
