package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá mundo!"
	canal <- "Teste"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
