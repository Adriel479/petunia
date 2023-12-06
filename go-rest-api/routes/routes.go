package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/", controllers.Home).Methods(http.MethodGet)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods(http.MethodGet)
	r.HandleFunc("/personalidades/{id}", controllers.Personalidade).Methods(http.MethodGet)
	r.HandleFunc("/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods(http.MethodDelete)
	r.HandleFunc("/personalidades/{id}", controllers.EditaUmaPersonalidade).Methods(http.MethodPut)
	r.HandleFunc("/personalidades", controllers.CriaUmaNovaPersonalidade).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8000", r))
}
