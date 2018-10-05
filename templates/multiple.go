package main

/*
https://golang.org/pkg/html/template/
https://gowebexamples.com/templates/
*/

import (
	"html/template"
	//"os"
)

import "math/rand"
import "time"
import "fmt"
import "bytes"

type TableRow struct {
	Name  string
	Count int
}

type Table struct {
	Title string
	Rows  []TableRow
}

type ReportHeader struct {
	MetaTitle       string
	MetaDescription string
}

type ReportFooter struct {
	Date string
}

func main() {
	dayst := new(bytes.Buffer)
	dayst = htmlHeader()
	fmt.Printf("%s", dayst)

	dayst = daysTable()
	fmt.Printf("%s", dayst)

	dayst = htmlFooter()
	fmt.Printf("%s", dayst)
	fmt.Printf("%s", "Done")
	//f, _ := os.Create("report.html")
}

func htmlFooter() *bytes.Buffer {
	tmpl := template.Must(template.ParseFiles("parts/footer.html"))
	buf := new(bytes.Buffer)

	currentTime := time.Now()
	currentTimeString := fmt.Sprintf(currentTime.Format("2006-01-02 15:04:05"))

	reportHeader := ReportFooter{Date: currentTimeString}

	tmpl.Execute(buf, reportHeader)
	return buf
}
func htmlHeader() *bytes.Buffer {
	tmpl := template.Must(template.ParseFiles("parts/header.html"))
	buf := new(bytes.Buffer)

	currentTime := time.Now()
	currentTimeString := fmt.Sprintf(currentTime.Format("2006-01-02 15:04:05"))

	reportHeader := ReportHeader{
		MetaTitle:       "Pipal Report",
		MetaDescription: fmt.Sprintf("Report generated on %s", currentTimeString),
	}

	tmpl.Execute(buf, reportHeader)
	return buf
}
func daysTable() *bytes.Buffer {
	tmpl := template.Must(template.ParseFiles("parts/table.html"))

	var day_names = []string{
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

	for _, day := range day_names {
		data[day] = random.Intn(50)
	}

	var days []TableRow

	for _, k := range day_names {
		//fmt.Println("Day:", k, "Count:", data[k])
		days = append(days, TableRow{Name: k, Count: data[k]})
	}

	pageData := Table{
		Title: "Days of the week",
		Rows:  days,
	}

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, pageData)
	return buf
}
