package routes

import (
	"log"
	"net/http"

	"github.com/William-Libero/go-rest-api/controllers"
	"github.com/William-Libero/go-rest-api/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleWare)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
