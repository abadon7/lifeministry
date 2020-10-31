package main

import (
	"encoding/json"
	//"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	//	"strconv"
)

func InitDataBase() error {
	err := db.MakeMigrations()
	if err != nil {
		panic(err)
	}
	return nil
}

func AddStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var newStudent Student
	err := decoder.Decode(&newStudent)
	if err != nil {
		http.Error(w, "error while parsing new student data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.SaveStudent(newStudent); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
	return
}

func GetStudents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	cellar, err := db.FindStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(cellar)
	return
}
