package models

type Livro struct {
	Id        int    `json:"id"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
	Src       string `json:"src"`
}

var Livros []Livro
