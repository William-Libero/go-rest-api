package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/William-Libero/go-rest-api/db"
	"github.com/William-Libero/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	fmt.Println(db.DB)
	db.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p []models.Personalidade
	db.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CriaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	db.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade []models.Personalidade
	db.DB.Delete(&personalidade, id)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	db.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	db.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
