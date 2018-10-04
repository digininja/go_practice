package main

/*
https://golang.org/pkg/html/template/
https://gowebexamples.com/templates/
*/

import (
	"html/template"
	"os"
)

import "math/rand"
import "time"

type DayData struct {
	Day   string
	Count int
}

type TodoPageData struct {
	PageTitle string
	Days      []DayData
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	// not crypto secure generateion of seed
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	data := TodoPageData{
		PageTitle: "Days of the week",
		Days: []DayData{
			{Day: "Monday", Count: random.Intn(50)},
			{Day: "Tuesday", Count: random.Intn(50)},
			{Day: "Wednesday", Count: random.Intn(50)},
			{Day: "Thursday", Count: random.Intn(50)},
			{Day: "Friday", Count: random.Intn(50)},
			{Day: "Saturday", Count: random.Intn(50)},
			{Day: "Sunday", Count: random.Intn(50)},
		},
	}
	f, _ := os.Create("report.html")

	tmpl.Execute(f, data)
}
