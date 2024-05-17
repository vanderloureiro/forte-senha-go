package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		password := generatePassword()
		response := map[string]string{"password": password}

		w.Header().Set("Content-Type", "application/json")
		json.NewEnconder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}

var words = [10]string{"feijao", "bicicleta", "churrasco", "sanduiche", "helicoptero", "geladeira", "ensolarado", "frigideira", "ornitorrinco", "intercambio"}

func generatePassword() string {
	randIndex := rand.Intn(10)
	return words[randIndex]
}
