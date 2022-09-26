package main

import "time"

type Student struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Gender      string    `json:"gender"`
	Cel         int       `json:"cel"`
	Active      bool      `json:"active"`
	Note        string    `json:"notes"`
	Last        time.Time `json:"last"`
	LastPartner int       `json:"lastpartner"`
}

type Couple struct {
	InCharge Student   `json:"incharge"`
	Helper   Student   `json:"helper"`
	Type     Assigment `json:"assigmenttype"`
	Date     time.Time `json:"date"`
}
