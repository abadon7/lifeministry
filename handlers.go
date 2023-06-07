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

func GetStudents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	active := r.URL.Query().Get("active")
	gender := r.URL.Query().Get("gender")

	fmt.Println("active = " + active)
	fmt.Println("gender = " + gender)

	if active == "" {
		active = "all"
	}

	if gender == "" {
		gender = "all"
	}

	cellar, err := db.FindStudents(active, gender)

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
	date := r.URL.Query().Get("date")

	couple, err := partnersmaker(num, gender, date)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(couple)
	return
}

func AddSchedule(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS,PUT")

	decoder := json.NewDecoder(r.Body)

	var newSchedule Schedule
	err := decoder.Decode(&newSchedule)
	if err != nil {
		http.Error(w, "error while parsing new Schedule data: "+err.Error(), http.StatusBadRequest)
		return
	}

	scheduleInfo, err := db.SaveSchedule(newSchedule)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scheduleInfo)
	return
}

func GetSchedules(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	cellar, err := db.FindSchedules()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(cellar)
	return
}

func GetSchedule(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid schedule id, it must be a number", ps.ByName("id")), http.StatusBadRequest)
		return
	}
	fmt.Println("Schedule # " + strconv.Itoa(ID))
	result, err := db.FindSchedule(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
	return
}

func GetScheduleToFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid schedule id, it must be a number", ps.ByName("id")), http.StatusBadRequest)
		return
	}
	fmt.Println("Schedule # " + strconv.Itoa(ID))
	result := generator(ID)
	//result := "new_result_1.docx"
	//	result, err := db.FindSchedule(ID)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//		return
	//	}
	//fileBytes, err := ioutil.ReadFile(result)
	//if err != nil {
	//	panic(err)
	//}
	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("Programa.docx"))
	w.Header().Set("Content-Type", "application/octet-stream")
	//w.Header().Add("Content-Description", "File Transfer")
	//w.Header().Add("Content-Type", "application/octet-stream")
	//w.Header().Add("Content-Transfer-Encoding", "binary")
	//w.Header().Add("Expires", "0")
	//w.Header().Add("Cache-Control", "must-revalidate")
	//w.Header().Add("Pragma", "public")
	//	w.Write(fileBytes)
	http.ServeFile(w, r, result)
	return
}

func GetS89Files(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid schedule id, it must be a number", ps.ByName("id")), http.StatusBadRequest)
		return
	}
	fmt.Println("Schedule # " + strconv.Itoa(ID))
	schedule, err := db.FindSchedule(ID)
	result := generateS89(schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
	return
}
