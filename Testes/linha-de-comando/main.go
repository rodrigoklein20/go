package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	for _, arg := range os.Args {
		log.Printf("%s", arg)
	}
	var entry string
	_, err := fmt.Scan(&entry)
	if err != nil {
		panic(err)
	}

	log.Printf("Usuário digitou: %s", entry)
}
