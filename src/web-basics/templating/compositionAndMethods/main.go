package main

import (
	"html/template"
	"log"
	"os"
)

var t *template.Template

type curso struct {
	Numero, Nome, Unidades string
}

type semestre struct {
	Term   string
	Cursos []curso
}

type ano struct {
	Inverno, Verao, Outono semestre
}

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	primeiroAno := ano{
		Inverno: semestre{
			Term: "Inverno",
			Cursos: []curso{
				curso{"40", "Introdução à matemática", "3"},
				curso{"80", "Introdução à arqueologia", "5"},
				curso{"32", "Introdução à biologia", "2"},
			},
		},
	}
	err := t.Execute(os.Stdout, primeiroAno)
	if err != nil {
		log.Fatal(err)
	}
}
