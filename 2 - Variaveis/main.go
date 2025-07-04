package main

import "fmt"

func main() {
	//Tipo explícito
	var variavel1 string = "Variável 1"
	//Tipo implícito
	variavel2 := "Variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Declarar diversas variáveis
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)
	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//Invertendo valores de variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
