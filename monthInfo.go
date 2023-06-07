package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

//func main() {
//	fmt.Println("vim-go")
//}

func getWeekNumber(date string) [2]int {
	fullDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
	fmt.Println(fullDate.Format(time.RFC3339))
	year, week := fullDate.ISOWeek()
	fmt.Println(year, week)
	return [2]int{week, year}
}

func getMonthInfo(dateKeys WeeksKeys) GroupWeekInfo {
	groupWeekInfo := make(GroupWeekInfo)
	//KeysData := [5]string{"2022-01-03T05:00:00Z", "2022-01-10T05:00:00Z", "2022-01-17T05:00:00Z", "2022-01-24T05:00:00Z", "2022-01-31T05:00:00Z"}
	//	keysData := dateKeys
	//numWeek := 20

	DateInfo := getWeekNumber(dateKeys[0])
	numWeek := DateInfo[0] - 1
	currentYear := DateInfo[1]
	//var weeksKeys WeeksKeys
	//weeksKeys = dateKeys
	// Find and visit all links
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("h1"))
	//})

	//for w := 0; w < 4; w++ {
	for index, key := range dateKeys {
		numWeek = numWeek + 1
		fmt.Println("Number of the Week", numWeek, index)
		c := colly.NewCollector()
		c.SetRequestTimeout(30 * time.Second)
		c.OnHTML("#content", func(e *colly.HTMLElement) {
			//date := e.ChildText("h1[data-pid=p1]")
			weekinfo := WeekInfo{}
			weekinfo.Date = e.ChildText("#p1")
			weekinfo.Text = e.ChildText("#p2")
			weekinfo.Song = e.ChildText(".bodyTxt #p3")
			weekinfo.Treasures = e.ChildText(".bodyTxt #p6")
			//fmt.Println(weekinfo.Date, weekinfo.Text)
			//fmt.Println(e.Text)

			e.ForEach("#section3 li ", func(_ int, el *colly.HTMLElement) {
				//fmt.Println("Econtrado", el.Text)
				weekinfo.School = append(weekinfo.School, el.Text)
			})
			e.ForEach("#section4 li ", func(_ int, el *colly.HTMLElement) {
				//fmt.Println("Econtrado", el.Text)
				weekinfo.Living = append(weekinfo.Living, el.Text)
			})
			//fmt.Println(e.ChildAttr("#footerNextWeek a", "href"))
			groupWeekInfo[key] = weekinfo
		})

		c.OnScraped(func(r *colly.Response) {
			fmt.Println("Finished", r.Request.URL)
			//	fmt.Println(weekinfo.Date, weekinfo.Text, weekinfo.Song, weekinfo.Treasures, weekinfo.School, weekinfo.Living)
			//fmt.Println(groupWeekInfo)
		})
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})
		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})

		c.Visit("https://wol.jw.org/es/wol/meetings/r4/lp-s/" + strconv.Itoa(currentYear) + "/" + strconv.Itoa(numWeek))
	}

	result, err := json.Marshal(groupWeekInfo)
	if err != nil {
		fmt.Printf("Error: %s ", err.Error)
	} else {
		fmt.Println(string(result))
	}

	return groupWeekInfo
}
