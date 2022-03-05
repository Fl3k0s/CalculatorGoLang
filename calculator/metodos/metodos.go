package metodos

import (
	"container/list"
	"fmt"
)

func Suma() {
	var suma int

	numeros := leerVariosNumeros()
	fmt.Println(numeros)
}

func Resta() {
	var resta int
	num1, num2 := leerDosNumeros("Cual es su numero a restar", "Por cuanto quiere restarlo")

	numeros := list.New()
	numeros.PushBack(num1)
	numeros.PushBack(num2)
}

func Multiplicacion() {
	var multiplicacion int

	numeros := leerVariosNumeros()
	fmt.Println(numeros)
}

func Division() {
	var division int
	num1, num2 := leerDosNumeros("Cual es el Dividendo", "Cual es el Divisor")

	numeros := list.New()
	numeros.PushBack(num1)
	numeros.PushBack(num2)
}

func Potencia() {
	var potencia int
	num1, num2 := leerDosNumeros("Cual es el numero", "Cual es el exponente")

	numeros := list.New()
	numeros.PushBack(num1)
	numeros.PushBack(num2)
}

func leerDosNumeros(mensaje1, mensaje2 string) (a int, b int) {
	//leer los 2 numeros de teclado
	fmt.Println(mensaje1)
	a = 0
	fmt.Println(mensaje2)
	b = 0
	return
}

func leerVariosNumeros() list.List {
	continuar := true
	numeros := list.New()
	count := 0
	for continuar {
		var numero int
		//leernumero
		numero = 1
		numeros.PushBack(numero)
		count++
		if count > 1 {
			fmt.Println("Quiere añadir otro numero?")
			//leer Y/N

			var respuesta string
			for respuesta != "Y" || respuesta != "y" || respuesta != "n" || respuesta != "N" {
				//leer respuesta
			}

			if respuesta != "Y" || respuesta != "y" {
				continuar = false
			}
		}
	}
	return *numeros
}

func llamadaApi() int {
	url := "http://localhost:8000"

}
