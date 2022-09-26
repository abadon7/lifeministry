package main

type Week map[string][]Couple

type Schedule struct {
	ID    int64  `json:"id"`
	Data  string `json:"data"`
	Range string `json:"range"`
}

type WeekInfo struct {
	Date      string   `json:"date"`
	Text      string   `json:"text"`
	Song      string   `json:"song"`
	Treasures string   `json:"treasures"`
	School    []string `json:"school"`
	Living    []string `json:"living"`
}

type GroupWeekInfo map[string]WeekInfo
type WeeksKeys []string
