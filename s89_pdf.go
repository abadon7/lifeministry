package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func s89ToPdf(data S89, name string, path string) {
	fmt.Println("Creating S89 file")
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Size")

	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set("A5")
	// 	htmlStr := `<html><body><h1 style="color:red;">This is an html
	//  from pdf to test color<h1><img src="http://api.qrserver.com/v1/create-qr-
	// code/?data=HelloWorld" alt="img" height="42" width="42"></img></body></html>`

	checkOff := "<input type='checkbox' id='checkbox-7c7a' name='checkbox' value='On' />"
	checkOn := "<input type='checkbox' id='checkbox-7c7a' name='checkbox' value='On' checked />"

	newData := data

	newData.CheckReading = checkOff
	newData.CheckFirst = checkOff
	newData.CheckStudy = checkOff
	newData.CheckReturn = checkOff
	newData.CheckTalk = checkOff

	fmt.Println("Checking Data")

	if data.CheckReading == "1" {
		newData.CheckReading = checkOn
	}
	if data.CheckFirst == "1" {
		newData.CheckFirst = checkOn
	}
	if data.CheckStudy == "1" {
		newData.CheckStudy = checkOn
	}
	if data.CheckReturn == "1" {
		fmt.Println("Return found")
		newData.CheckReturn = checkOn
	}
	fmt.Println(newData.CheckReturn)
	if data.CheckTalk == "1" {
		newData.CheckTalk = checkOn
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}

	// Render the HTML from the template
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, newData)
	if err != nil {
		log.Fatal(err)
	}
	htmlStr := buffer.String()

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))

	fmt.Println("Creating Buffer")

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile(path + name + ".pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
