package main

import (
	"fmt"
)

func main() {
	fmt.Println("Calculadora main")
}

// esto será la funcion que muestre el menú y llame al menú
func menu() {
	fmt.Printf("Que operación desea realizar:")
	fmt.Println("1. Suma")
	fmt.Println("2. Resta")
	fmt.Println("3. Multiplicacion")
	fmt.Println("4. Division")
	fmt.Println("5. Potencia")
	fmt.Println("6. Salir")

	var number int
	continues := false

	for continues != true {
		//introducir metodo para leer numero de teclado
		number = 1
		if number > 0 && number < 7 {
			continues = true
		}
	}

	funcionalidades(number)
}

func funcionalidades(n int) {
	switch n {
	case 1:
		//metodo de suma
		break
	case 2:
		//metodo de resta
		break
	case 3:
		//metodo de multiplicacion
		break
	case 4:
		//metodo de division
		break
	case 5:
		//metodo de potencia
		break
	case 6:
		//exit
		break
	}
}
