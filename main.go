package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w, "Home Page")

}

func HandleRequest() {

	http.HandleFunc("/", Home)

}

func main() {

	fmt.Println("Iniciando o servidor Rest com Go.")
	HandleRequest()
	log.Fatal(http.ListenAndServe(":8000", nil))

}