package routes

import (
	"log"
	"net/http"

	"github.com/tiagosilva-dev/react-books-backend/controllers"
)

func HandleRequest() {

	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/livros", controllers.TodosLivros)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
