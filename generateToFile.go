package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/nguyenthenguyen/docx"
)

func generateFile(schedules []Schedule, scheduleKeys []string, monthInfo GroupWeekInfo) string {
	// Read from docx file
	r, err := docx.ReadDocxFile("./S-140-S_3.docx")
	fileName := "./new_result_1.docx"
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)
	if err != nil {
		panic(err)
	}

	meses := [12]string{"ENERO", "FEBRERO", "MARZO", "ABRIL", "MAYO", "JUNIO", "JULIO", "AGOSTO", "SEPTIEMBRE", "OCTUBRE", "NOVIEMBRE", "DICIEMBRE"}
	docx1 := r.Editable()
	// Replace like https://golang.org/pkg/strings/#Replace
	//	docx1.Replace("[Asigna1]", "Henry 1", -1)
	//	docx1.Replace("[Asigna2]", "Henry 2", -1)
	//	docx1.Replace("[FECHA1]", "Nueva fecha", -1)
	docx1.Replace("[NOMBRE DE LA CONGREGACIÓN]", "NECHÍ", -1)
	//JSON exapmple
	//filePath := "./data.json"
	//fmt.Printf("// reading file %s\n", filePath)
	//file, err1 := ioutil.ReadFile(filePath)
	//if err1 != nil {

	//	fmt.Printf("// error reading file %s\n", filePath)

	//	fmt.Printf("File error: %v\n", err1)
	//	os.Exit(1)
	//}
	//fmt.Println("// defining array of struct schedules")
	//var schedules []Schedule

	//err2 := json.Unmarshal(file, &schedules)
	//if err2 != nil {
	//	fmt.Println("error:", err2)
	//	os.Exit(1)
	//}

	SchDate := 0
	Asig := 0

	var WDateLabel string
	fmt.Println("// loop over array of structs of schedules")
	for s := range schedules {
		//var weeks Week
		fmt.Printf("The ship '%d' first appeared on '%v'\n", schedules[s].ID, schedules[s].Range)
		//scheduleKeys := getScheduleKeys(schedules[s].Data)
		//scheduleKeys := make([]string, 0, len(schedules[s].Data))
		//	for sk := range schedules[s].Data {
		//		scheduleKeys = append(scheduleKeys, sk)
		//	}
		//	fmt.Println(scheduleKeys)
		//	sort.Strings(scheduleKeys)
		//monthInfo := getMonthInfo(scheduleKeys)
		//	fmt.Println(scheduleKeys)

		//for k, ws := range schedules[s].Data {
		for k, ws := range scheduleKeys {
			fmt.Printf("Key: '%d' : Data: '%v'\n", k, ws)
			//SchDate++
			SchDate = k + 1
			weekRange, err := time.Parse(time.RFC3339, ws)
			if err != nil {
				log.Fatal(err)
			}
			fstDay := weekRange.Day()
			fnlDay := weekRange.AddDate(0, 0, 6)
			WDateLabel = strconv.Itoa(fstDay) + " A " + strconv.Itoa(fnlDay.Day()) + " DE " + meses[int(fnlDay.Month())-1]
			if fnlDay.Day() < fstDay {
				WDateLabel = strconv.Itoa(fstDay) + " DE " + meses[int(weekRange.Month())-1] + " A " + strconv.Itoa(fnlDay.Day()) + " DE " + meses[int(fnlDay.Month())-1]
			}
			fmt.Println(weekRange.Day(), weekRange.AddDate(0, 0, 6).Day())
			fmt.Printf("start: '%v' '%v' : end: '%v' '%v'\n", weekRange.Day(), weekRange.Month(), weekRange.AddDate(0, 0, 6).Day(), weekRange.AddDate(0, 0, 6).Month())
			currentDate := "[FECHA" + strconv.Itoa(SchDate) + "]"
			currentText := "[LECTURA" + strconv.Itoa(SchDate) + "]"
			currentTreasures := "[TESOROS" + strconv.Itoa(SchDate) + "]"
			currentSong1 := "[Número" + strconv.Itoa(SchDate) + "]"
			fmt.Println(currentDate)
			docx1.Replace(currentDate, WDateLabel, -1)
			fmt.Println(currentText, monthInfo[ws].Text)
			docx1.Replace(currentText, monthInfo[ws].Text, -1)
			fmt.Println(currentTreasures, monthInfo[ws].Treasures)
			docx1.Replace(currentTreasures, monthInfo[ws].Treasures, -1)
			docx1.Replace(currentSong1, monthInfo[ws].Song, -1)

			var weekSchedules Week
			errw := json.Unmarshal([]byte(schedules[s].Data), &weekSchedules)
			if errw != nil {
				fmt.Println("error:", err)
			}

			for kw, w := range weekSchedules[ws] {
				fmt.Println(kw)
				CurrentAssigNumber := (k * 4) + kw + 1
				fmt.Println(CurrentAssigNumber)
				fmt.Println(w.InCharge.Name)
				Asig++
				ParticipantsNames := w.InCharge.Name
				if w.Helper.ID != 0 {
					ParticipantsNames = w.InCharge.Name + "/" + w.Helper.Name
				}
				assigNumber := "[Asigna" + strconv.Itoa(CurrentAssigNumber) + "]"
				fmt.Println(assigNumber)
				docx1.Replace(assigNumber, ParticipantsNames, 1)
			}
		}
		docx1.WriteToFile(fileName)
		r.Close()
		//for k, v := range schedules[s].Data.string {
		//	fmt.Printf("Key: %d : | Value : %v\n", k, v)
		//}

		//		err3 := json.Unmarshal(schedules[s].Data, &weeks)
		//		if err3 != nil {
		//			fmt.Println("error:", err3)
		//			os.Exit(1)
		//		}
		//		fmt.Printf("'%v'\n", weeks)
	}
	return fileName
}
