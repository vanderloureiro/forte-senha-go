package main

import (
	"fmt"
	"net/http"
	"math/rand"
)

func main() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, generatePassword())
	})

	http.ListenAndServe(":8080", nil)
}

var words = [10]string{"feijao", "bicicleta", "churrasco", "sanduiche", "helicoptero", "geladeira", "ensolarado", "frigideira", "ornitorrinco", "intercambio"}

func generatePassword() string {
	randIndex := rand.Intn(10)
	return words[randIndex]
}
