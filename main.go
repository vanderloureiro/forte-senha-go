package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", h1)

	http.ListenAndServe(":8000", nil)
}

var words = [10]string{"feijao", "bicicleta", "churrasco", "sanduiche", "helicoptero", "geladeira", "ensolarado", "frigideira", "ornitorrinco", "intercambio"}

func generatePassword() string {
	randIndex := rand.Intn(10)
	return words[randIndex]
}
