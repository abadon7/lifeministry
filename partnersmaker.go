package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Couple struct {
	InCharge Student   `json:"incharge"`
	Helper   Student   `json:"helper"`
	Type     Assigment `json:"assigmenttype"`
	Date     time.Time `json:"date"`
}

func partnersmaker(tp int, gender string) (Couple, error) {
	var err error
	var couple Couple
	var randnum2 int
	var helper Student
	today := time.Now()

	students, err := db.FindStudents("active", gender)
	if err != nil {
		panic(err)
	}

	assigment, err := db.FindAssigment(Assigment{ID: tp})
	if err != nil {
		return couple, err
	}

	randnum := rand.Intn(len(students))
	incharge := students[randnum]

	if assigment[0].Participants > 1 {
		randnum2 = rand.Intn(len(students))
		helper = students[randnum2]
		for randnum == randnum2 || incharge.LastPartner == helper.ID {
			fmt.Println("This ID is no available" + strconv.Itoa(randnum2))
			randnum2 = rand.Intn(len(students))
			helper = students[randnum2]
		}

	}

	couple.InCharge = incharge
	couple.Helper = helper
	couple.Date = today
	couple.Type = assigment[0]

	incharge.Last = today
	incharge.LastPartner = helper.ID

	helper.Last = today
	helper.LastPartner = incharge.ID

	//if err := db.UpdateStudent(incharge); err != nil {
	//	return couple, err
	//}

	//if err := db.UpdateStudent(helper); err != nil {
	//	return couple, err
	//}

	return couple, nil
}
