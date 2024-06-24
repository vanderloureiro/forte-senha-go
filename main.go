package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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
	word = strings.Replace(word, "a", "@", 1)
	word = strings.Replace(word, "e", "&", 1)
	if len(word)%2 == 0 {
		old := string(word[0])
		new := strings.ToUpper(old)
		word = strings.Replace(word, old, new, 1)
	}
	word = word + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10))
	return word
}
