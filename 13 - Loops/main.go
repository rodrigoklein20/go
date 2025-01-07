package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("I:", i)
		//time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("J:", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"joão", "davi", "lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for i, letra := range "PALAVRA" {
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
