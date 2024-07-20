package routes

import (
	"api-rest-celebridades/controllers"
	"api-rest-celebridades/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidadePorId).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.AtualizaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8080", r))
}
