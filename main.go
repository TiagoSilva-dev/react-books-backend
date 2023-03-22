package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tiagosilva-dev/react-books-backend/routes"
)

func main() {

	log.Printf("Connect to http://localhost:4010/api/livros")
	log.Printf("[!] To exit press CTRL+C \n")

	routes.HandleRequest()
}
