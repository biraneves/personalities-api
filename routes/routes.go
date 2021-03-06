package routes

import (
	"log"
	"net/http"

	"github.com/biraneves/personalities-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.RetrievePersonality).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))

}
