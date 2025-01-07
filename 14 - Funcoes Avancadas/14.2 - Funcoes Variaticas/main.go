package main

import "fmt"

//O parâmetro é recebido na função como um slice, então é possível iterar sobre eles usando o range
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	resultado := soma(10, 20, 15)
	fmt.Println(resultado)

	escrever("Olá mundo", 10, 2, 3, 5)
}
