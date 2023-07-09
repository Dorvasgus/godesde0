package ejercicios

import "strconv"

func Numerocien(valor string) (int, string) {

	if numero, err := strconv.Atoi(valor); numero > 100 || err == nil {
		return numero, "El numero es mayor a 100"
	} else {
		return numero, "EL numero es menor a 100"
	}
}
