package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tiagosilva-dev/react-books-backend/controllers"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/livros", controllers.TodosLivros).Methods("Get")
	r.HandleFunc("/api/livros/{id}", controllers.RetornaUmLivro).Methods("Get")

	r.Use(mux.CORSMethodMiddleware(r))

	log.Fatal(http.ListenAndServe(":4010", r))
}
