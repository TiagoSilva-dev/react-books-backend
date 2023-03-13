package main

import (
	"fmt"

	"github.com/tiagosilva-dev/react-books-backend/models"
	"github.com/tiagosilva-dev/react-books-backend/routes"
)

func main() {

	models.Livros = []models.Livro{
		{Nome: "Nome 1", Descricao: "Descricao 1"},
		{Nome: "Nome 2", Descricao: "Descricao 2"},
	}
	fmt.Println("Iniciando o Servidor Rest com go")
	routes.HandleRequest()
}
