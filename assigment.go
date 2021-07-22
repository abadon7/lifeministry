package main

type Assigment struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Participants int    `json:"participants"`
}

type Schedule struct {
	ID    int64  `json:"id"`
	Data  string `json:"data"`
	Range string `json:"range"`
}
