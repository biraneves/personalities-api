package main

import (
	"fmt"

	"github.com/biraneves/personalities-api/models"
	"github.com/biraneves/personalities-api/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{Name: "Name 1", History: "History 1"},
		{Name: "Name 2", History: "History 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go.")
	routes.HandleRequest()

}
