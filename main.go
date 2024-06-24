package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

type Password struct {
	PasswordValue string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		if r.Method == "POST" {
			password := Password{generatePassword()}
			tmpl.Execute(w, password)
		} else {
			tmpl.Execute(w, nil)
		}

	}
	http.HandleFunc("/", h1)

	http.ListenAndServe(":8000", nil)
}

var words = [10]string{"feijao", "bicicleta", "churrasco", "sanduiche", "helicoptero", "geladeira", "ensolarado", "frigideira", "ornitorrinco", "intercambio"}

func generatePassword() string {
	randIndex := rand.Intn(10)
	word := words[randIndex]

	runes := []rune(word)

	for i := 0; i < len(word); i++ {
		if word[i] == 'a' {
			runes[i] = '@'
		}
	}
	return string(runes)
}
