package routes

import (
	"net/http"

	"github.com/biraneves/personalities-api/controllers"
)

func HandleRequest() {

	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.AllPersonalities)

}
