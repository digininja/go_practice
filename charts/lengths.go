package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

import "math/rand"
import "time"
import "strconv"

// Used in the plotter colours
const RED = 0
const GREEN = 1
const BLUE = 2
const WIDTH = 5 * vg.Inch
const HEIGHT = 3 * vg.Inch

func GenBarChart(filename string, title string, ylabel string, xlabels []string, data []float64) {
	groupAY := plotter.Values{}

	for _, val := range data {
		groupAY = append(groupAY, float64(val))
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = title
	p.Y.Label.Text = ylabel

	// Default font size
	w := vg.Points(50)

	barsA, err := plotter.NewBarChart(groupAY, w)
	if err != nil {
		panic(err)
	}
	// barsA.LineStyle.Width = vg.Length(0)
	// barsA.Offset = -w
	barsA.Color = plotutil.Color(BLUE)
	barsA.Width = 20

	p.Add(barsA)
	// p.Legend.Add("Group A", barsA)
	p.NominalX(xlabels...)

	if err := p.Save(WIDTH, HEIGHT, filename); err != nil {
		panic(err)
	}

}

func main() {
	var lengths = make([]float64, 10)

	// not crypto secure generateion of seed
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	for i := 0; i < 10; i++ {
		lengths[i] = float64(random.Intn(50))
	}

	var xlabels = make([]string, len(lengths))

	i := 0
	for k := range lengths {
		xlabels[i] = strconv.Itoa(k)
		i++
	}

	GenBarChart("lengths.png", "Password Length", "Number of characters", xlabels, lengths)
}
