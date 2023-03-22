package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tiagosilva-dev/react-books-backend/internal/db"
	"github.com/tiagosilva-dev/react-books-backend/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func TodosLivros(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	dbcon, err := sql.Open("mysql", "root:supplylabs@tcp(localhost:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}

	defer dbcon.Close()

	dt := db.New(dbcon)

	livros, _ := dt.GetLivros(ctx)

	json.NewEncoder(w).Encode(livros)
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
