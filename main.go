package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/biraneves/personalities-api/routes"
)

func main() {

	fmt.Println("Iniciando o servidor Rest com Go.")
	routes.HandleRequest()
	log.Fatal(http.ListenAndServe(":8000", nil))

}
