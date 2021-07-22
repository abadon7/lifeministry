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

func partnersmaker(tp int, gender string, date string) (Couple, error) {
	var err error
	var couple Couple
	var randnum int
	var randnum2 int
	var incharge Student
	var helper Student
	var inlast float64
	var hlplast float64
	var intry int
	var hlptry int

	fDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		panic(err)
	}
	//fDate := time.Now()

	students, err := db.FindStudents("active", gender)
	if err != nil {
		panic(err)
	}

	assigment, err := db.FindAssigment(Assigment{ID: tp})
	if err != nil {
		return couple, err
	}

	randnum = rand.Intn(len(students))
	incharge = students[randnum]
	//	today:= time.Now()
	inlast = fDate.Sub(incharge.Last).Hours() / 24

	for inlast <= 45 && intry < 10 {
		randnum = rand.Intn(len(students))
		incharge = students[randnum]
		inlast = fDate.Sub(incharge.Last).Hours() / 24
		intry++
	}

	if assigment[0].Participants > 1 {
		randnum2 = rand.Intn(len(students))
		helper = students[randnum2]
		hlplast = fDate.Sub(helper.Last).Hours() / 24
		fmt.Println(hlplast)
		//		if hlptry < 10 {
		for (randnum == randnum2 || incharge.LastPartner == helper.ID || hlplast <= 45) && (hlptry < 10) {
			fmt.Println("This ID is no available" + strconv.Itoa(randnum2))
			randnum2 = rand.Intn(len(students))
			helper = students[randnum2]
			hlplast = fDate.Sub(helper.Last).Hours() / 24
			hlptry++
		}

	}
	//	}

	incharge.Last = fDate
	incharge.LastPartner = helper.ID

	helper.Last = fDate
	helper.LastPartner = incharge.ID

	couple.InCharge = incharge
	couple.Helper = helper
	couple.Date = fDate
	couple.Type = assigment[0]

	if err := db.UpdateStudent(incharge); err != nil {
		return couple, err
	}

	if helper.ID > 0 {
		if err := db.UpdateStudent(helper); err != nil {
			return couple, err
		}
	}

	return couple, nil
}
