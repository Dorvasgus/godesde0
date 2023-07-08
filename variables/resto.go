package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float64
var Fecha time.Time

func Restovariable() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1522.12
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func Convietyoatexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)

	return true, texto
}
