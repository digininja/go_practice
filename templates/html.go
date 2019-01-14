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
import "fmt"

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

	var days = []string{
		"monday",
		"tuesday",
		"wednesday",
		"thursday",
		"friday",
		"saturday",
		"sunday",
	}

	var data = make(map[string]int)

	// not crypto secure generation of seed
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	for _, day := range days {
		data[day] = random.Intn(50)
	}

	var aDays []DayData

	for _, k := range days {
		fmt.Println("Day:", k, "Value:", data[k])
		aDays = append(aDays, DayData{Day: k, Count: data[k]})
	}

	pageData := TodoPageData{
		PageTitle: "Days of the week",
		Days:      aDays,
	}
	f, _ := os.Create("report.html")

	tmpl.Execute(f, pageData)
}
