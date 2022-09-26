package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func getScheduleKeys(schedule string) []string {
	var sch Week
	err := json.Unmarshal([]byte(schedule), &sch)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	scheduleKeys := make([]string, 0, len(sch))
	for sk := range sch {
		scheduleKeys = append(scheduleKeys, sk)
	}
	fmt.Println(scheduleKeys)
	sort.Strings(scheduleKeys)
	return scheduleKeys
}

func generator() {
	filePath := "./data.json"
	fmt.Printf("// reading file %s\n", filePath)
	file, err1 := ioutil.ReadFile(filePath)
	if err1 != nil {

		fmt.Printf("// error reading file %s\n", filePath)

		fmt.Printf("File error: %v\n", err1)
		os.Exit(1)
	}
	fmt.Println("// defining array of struct schedules")
	var schedules []Schedule

	err2 := json.Unmarshal(file, &schedules)
	if err2 != nil {
		fmt.Println("error:", err2)
		os.Exit(1)
	}

	scheduleKeys := getScheduleKeys(schedules[0].Data)
	monthInfo := getMonthInfo(scheduleKeys)
	generateFile(schedules, scheduleKeys, monthInfo)
}
