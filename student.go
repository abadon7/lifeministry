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
