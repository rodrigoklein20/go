package main

import (
	"fmt"
	"strconv"
)

func twoSums(nums []int, target int) []int {
	// Example 1:

	// Input: nums = [2,7,11,15], target = 9
	// Output: [0,1]
	// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

	fmt.Printf("Resultado esperado: %d \n", target)
	var soma int
	arrayRetorno := []int{}
	for i, numeroI := range nums {
		for j, numeroJ := range nums {
			if i != j {
				soma = numeroI + numeroJ
				fmt.Printf("Posição i: - %d Posição j: %d - Soma = %d \n", i, j, soma)
				if soma == target {
					arrayRetorno = append(arrayRetorno, i)
					arrayRetorno = append(arrayRetorno, j)
					break
				}
			}
		}
		if len(arrayRetorno) >= 1 {
			break
		}
	}
	fmt.Println(arrayRetorno)
	return arrayRetorno
}

func isPalindrome(x int) bool {
	// Given an integer x, return true if x is a palindrome, and false otherwise.

	//Converte o int para uma string
	strNum := strconv.Itoa(x)

	//Itera sobre a string
	for p, digito := range strNum {
		//Busca o índice da posição inversa a atual \ length - (p+1)
		pInversa := len(strNum) - (p + 1)
		//Se o digito atual for diferente do digito da posição inversa retorna false
		if string(digito) != string(strNum[pInversa]) {
			return false
		}
	}
	return true
}

func main() {
	//Desafio 1
	//array := []int{0, 4, 3, 0}
	//fmt.Println(twoSums(array, 0))

	//Desafio 2
	fmt.Println(isPalindrome(-18855881))
}
