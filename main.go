package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

//const ServerAddr = ":8080"

var db Storage
var router *httprouter.Router

func init() {
	var err error

	db, err = NewStorage(SQLite)
	if err != nil {
		log.Fatal(err)
	}

	err = InitDataBase()
	if err != nil {
		log.Fatal(err)
	}

	router = httprouter.New()
	router.GET("/", welcome)
	router.GET("/students", GetStudents)
	router.GET("/students/:id", GetStudent)
	router.GET("/assignments", GetAssigments)
	router.GET("/partners", GetPartners)
	router.GET("/schedules", GetSchedules)
	router.GET("/schedules/:id", GetSchedule)
	router.GET("/getschedule/:id", GetScheduleToFile)
	router.GET("/gets89/:id", GetS89Files)

	router.POST("/students", AddStudent)
	router.POST("/assignments", AddAssigment)
	router.POST("/schedules", AddSchedule)

	router.PUT("/students", UpdtStudent)

	router.DELETE("/students/:id", DeleteStudent)

}

func welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to the System")
}

func getListenPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	port, err := getListenPort()
	if err != nil {
		log.Fatal(err)
	}
	_cors := cors.Options{
		AllowedMethods: []string{"POST", "PUT", "OPTIONS", "GET"},
		AllowedOrigins: []string{"*"},
	}
	handler := cors.New(_cors).Handler(router)

	fmt.Println("The server is on port " + port)
	log.Fatal(http.ListenAndServe(port, handler))
	//log.Fatal(http.ListenAndServe(port, router))
}
