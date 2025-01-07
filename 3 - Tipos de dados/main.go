package main

import (
	"errors"
	"fmt"
)

func main() {

	//INTEIROS - INICIO
	//int8, int16, int32, int64
	var numero int = 10000000000
	fmt.Println(numero)

	// unsigned int - UINT
	var numero2 uint = 1000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//UINT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)
	//INTEIROS - FIM

	//REAIS - INICIO
	var numeroReal1 float32 = 123.23
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 121312323.23
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.42
	fmt.Println(numeroReal3)
	//REAIS - FIM

	//STRINGS - INICIO
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)
	//STRINGS - FIM

	//Valor zero
	var texto string
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
