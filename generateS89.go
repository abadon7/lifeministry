package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func generateS89(studentsWeekInfo Schedule) string {
	var weekSchedules Week
	errw := json.Unmarshal([]byte(studentsWeekInfo.Data), &weekSchedules)
	if errw != nil {
		fmt.Println("error:", errw)
	}

	currentDate, err := time.Parse("1/2/2006", strings.Split(studentsWeekInfo.Range, "-")[0])
	if err != nil {
		log.Fatal(err)
	}

	year := currentDate.Year()
	month := currentDate.Month().String()
	path := "./output/" + strconv.Itoa(year) + "/" + month + "/"

	errw = os.MkdirAll(path, 0755)
	if errw != nil {
		log.Fatal(errw)
	}

	for _, wInfo := range weekSchedules {
		for _, aInfo := range wInfo {
			data := S89{
				Name:         aInfo.InCharge.Name,
				Helper:       aInfo.Helper.Name,
				Date:         aInfo.Date.Format("Jan 2, 2006"),
				CheckReturn:  "0",
				CheckStudy:   "0",
				CheckFirst:   "0",
				CheckTalk:    "0",
				CheckReading: "0",
			}
			if aInfo.Type.ID == 1 {
				data.CheckReading = "1"
			}
			if aInfo.Type.ID == 2 {
				data.CheckTalk = "1"
			}
			if aInfo.Type.ID == 3 {
				data.CheckReturn = "1"
			}
			if aInfo.Type.ID == 4 {
				data.CheckStudy = "1"
			}
			if aInfo.Type.ID == 5 {
				data.CheckFirst = "1"
			}
			fmt.Println(aInfo.InCharge.Name)
			fmt.Println(data.CheckReturn)
			fmt.Println(data.Date)
			s89ToPdf(data, aInfo.InCharge.Name, path)
		}

	}
	return "S89"
}
