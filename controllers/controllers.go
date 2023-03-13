package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tiagosilva-dev/react-books-backend/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func TodosLivros(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Livros)
}
