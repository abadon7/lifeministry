package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Couple struct {
	InCharge Student `json:"incharge"`
	Helper   Student `json:"helper"`
}

func partnersmaker(num int, gender string) (Couple, error) {
	var err error
	var couple Couple
	var randnum2 int
	students, err := db.FindStudents("active", gender)
	if err != nil {
		panic(err)
	}

	randnum := rand.Intn(len(students))
	randnum2 = rand.Intn(len(students))

	for randnum == randnum2 {
		fmt.Println("This number is used " + strconv.Itoa(randnum2))
		randnum2 = rand.Intn(len(students))
	}

	couple.InCharge = students[randnum]
	couple.Helper = students[randnum2]

	return couple, err
}
