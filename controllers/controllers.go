package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tiagosilva-dev/react-books-backend/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func TodosLivros(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Livros)
}

func RetornaUmLivro(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	for _, livro := range models.Livros {
		if strconv.Itoa(livro.Id) == id {
			json.NewEncoder(w).Encode(livro)
		}
	}
}
