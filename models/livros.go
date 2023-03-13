package models

type Livro struct {
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
}

var Livros []Livro
