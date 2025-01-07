package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Rodrigo"
	u.idade = 20
	fmt.Println(u)

	enderecoExemplo := endereco{"Tapes", 6}

	usuario2 := usuario{"Rodrigo", 20, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Rodrigo"}
	fmt.Println(usuario3)
}
