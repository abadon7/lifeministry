package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
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

	InitDataBase()

	router = httprouter.New()
	router.GET("/", welcome)
	router.GET("/students", GetStudents)
	router.GET("/students/:id", GetStudent)
	router.GET("/assigments", GetAssigments)
	router.GET("/partners", GetPartners)

	router.POST("/students", AddStudent)
	router.POST("/assigments", AddAssigment)

	router.PUT("/students", UpdtStudent)

	router.DELETE("/students/:id", DeleteStudent)

}

func welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to the system")
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
	fmt.Println("The server is on port " + port)
	log.Fatal(http.ListenAndServe(port, router))
}
