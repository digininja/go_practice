package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

import "math/rand"
import "fmt"
import "time"

func main() {
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

	// not crypto secure generateion of seed
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	for _, day := range days {
		data[day] = random.Intn(50)
	}

	groupAY := plotter.Values{}
	for _, k := range days {
		fmt.Println("Day:", k, "Value:", data[k])
		groupAY = append(groupAY, float64(data[k]))
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Days of the week"
	p.Y.Label.Text = "Number of instances"

	w := vg.Points(10)

	barsA, err := plotter.NewBarChart(groupAY, w)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = -w

	p.Add(barsA)
	//p.Legend.Add("Group A", barsA)
	p.NominalX(days...)

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "days.png"); err != nil {
		panic(err)
	}
}
