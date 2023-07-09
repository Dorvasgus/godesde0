package main

import (
	"dorvasgus/godesde0/ejercicios"
	// "dorvasgus/godesde0/variables"
	"fmt"
	// "runtime"
)

func main() {
	// estado, texto := variables.Convietyoatexto(1500)
	// fmt.Println(estado)
	// fmt.Println(texto)

	numero, cadena := ejercicios.Numerocien("500")
	fmt.Println(numero)
	fmt.Println(cadena)

	// if os := runtime.GOOS; os == "Linux." || os == "windows" {
	// 	fmt.Println("El sistema operativo es", os)
	// }

	// switch os := runtime.GOOS; os {
	// case "windows":
	// 	fmt.Println("Esto es windows")
	// case "darwin":
	// 	fmt.Println("Esto es darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }
}
