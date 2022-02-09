// Correção: 0,5. Você faz dois tesets, sendo que um com certeza irá falhar. O problema é que as notas dos alunos deveria ser uma fatia de inteiros, logo nunca irá conseguir
// deserializar, seja por Disciplina ou por Aluno.
package main

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"html/template"
	"os"
)

type DisciplinaPage struct {
	PageTitle string
	Turma []Disciplina
}

type Disciplina struct {
	Titulo string
	Professor string
	Semestre string
	Alunos []Aluno
}

type Aluno struct {
	Nome string
	N1 float32
	N2 float32
	N3 float32
	Media float32
	Situacao bool
}

func main() {
	var turma []Disciplina
	var notas []Aluno

	file, err := ioutil.ReadFile ("notas.json")

	if err != nil {
		log.Fatal("JSON read from file failed: ", err)
	}

	if err = json.Unmarshal([]byte(file), &turma); err != nil {
		log.Fatal("JSON unmarshaling failed: ", err)
	}

	if err = json.Unmarshal([]byte(file), &notas); err != nil {
		log.Fatal("JSON unmarshaling failed: ", err)
	}

	data:= DisciplinaPage {
		PageTitle: "",
		Turma: turma,
	}

	data2:= Disciplina {
		Alunos: notas,
	}

	t := template.Must(template.ParseFiles("notas.tmpl"))

	if err := t.Execute (os.Stdout, data); err != nil {
		log.Fatal(err)
	}
	if err := t.Execute (os.Stdout, data2); err != nil {
		log.Fatal(err)
	}

}