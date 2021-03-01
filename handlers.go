package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func InitDataBase() error {
	err := db.MakeMigrations()
	if err != nil {
		panic(err)
	}
	return nil
}

func AddStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS,PUT")

	decoder := json.NewDecoder(r.Body)

	var newStudent []Student
	err := decoder.Decode(&newStudent)
	if err != nil {
		http.Error(w, "error while parsing new student data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.SaveStudent(newStudent...); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
	return
}

func GetStudents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cellar, err := db.FindStudents("all", "all")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(cellar)
	return
}

func GetStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid student id, it must be a number", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	result, err := db.FindStudent(Student{ID: ID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return
}

func UpdtStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	decoder := json.NewDecoder(r.Body)

	var newStudent Student
	err := decoder.Decode(&newStudent)
	if err != nil {
		http.Error(w, "error while parsing update of student data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.UpdateStudent(newStudent); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
	return
}

func DeleteStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid student id, it must be a number", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	if err := db.RemoveStudent(Student{ID: ID}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
	return
}

func AddAssigment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS,PUT")

	decoder := json.NewDecoder(r.Body)

	var newAssigment []Assigment
	err := decoder.Decode(&newAssigment)
	if err != nil {
		http.Error(w, "error while parsing new assigment data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.SaveAssigment(newAssigment...); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
	return
}

func GetAssigments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	cellar, err := db.FindAssigments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(cellar)
	return
}

func GetPartners(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	num, err := strconv.Atoi(r.URL.Query().Get("num"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	gender := r.URL.Query().Get("gender")

	couple, err := partnersmaker(num, gender)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(couple)
	return
}
