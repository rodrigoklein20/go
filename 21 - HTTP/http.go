package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // raiz
		w.Write([]byte("Página Raiz!"))
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar página de usuários!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
