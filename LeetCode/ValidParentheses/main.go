package main

import (
	"log"
	"slices"
	"strings"
)

func main() {

	str := "()"
	log.Printf("%s é valido? %t\n\n", str, isValid(str))

	str = "()[]{}"
	log.Printf("%s é valido? %t\n\n", str, isValid(str))

	str = "(]"
	log.Printf("%s é valido? %t\n\n", str, isValid(str))

	str = "([{}])"
	log.Printf("%s é valido? %t\n\n", str, isValid(str))

	str = "(){}}{"
	log.Printf("%s é valido? %t\n\n", str, isValid(str))

	str = "["
	log.Printf("%s é valido? %t\n\n", str, isValid(str))

	str = "(([]){})"
	log.Printf("%s é valido? %t\n\n", str, isValid(str))

}

func isValid(s string) bool {

	dicionario := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	arrayDaString := strings.Split(s, "")

	tam := len(arrayDaString)
	log.Printf("Tamanho inicial %d", tam)
	retorno := true

	if tam%2 != 0 {
		return false
	}

	for i := 0; i < tam-1; i++ {
		if dicionario[arrayDaString[i]] == arrayDaString[i+1] {
			log.Printf("L - Removendo %s e %s", arrayDaString[i], arrayDaString[i+1])
			arrayDaString = slices.Delete(arrayDaString, i, i+2)

			tam -= 2
			i--
			log.Printf("Tamanho após iteração %d", tam)
		} else if dicionario[arrayDaString[i]] == arrayDaString[tam-(i+1)] {
			log.Printf("I - Removendo %s e %s", arrayDaString[i], arrayDaString[tam-(i+1)])

			arrayDaString = slices.Delete(arrayDaString, i, i+1)
			tam -= 1
			arrayDaString = slices.Delete(arrayDaString, tam-(i+1), tam-i)
			tam -= 1
			i--
			log.Printf("Tamanho após iteração %d", tam)
		} else {
			retorno = false
		}
	}

	return retorno

}
