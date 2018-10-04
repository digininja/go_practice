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
import "strconv"

func main() {
	var lengths = [10]int{}

	// not crypto secure generateion of seed
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	for i := 0; i < 10; i++ {
		lengths[i] = random.Intn(50)
	}

	groupAY := plotter.Values{}
	for length, count := range lengths {
		fmt.Println("Length:", length, "Value:", count)
		groupAY = append(groupAY, float64(count))
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Password Length"
	p.Y.Label.Text = "Number of characters"

	w := vg.Points(10)

	barsA, err := plotter.NewBarChart(groupAY, w)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = -w

	keys := make([]string, len(lengths))

	i := 0
	for k := range lengths {
		keys[i] = strconv.Itoa(k)
		i++
	}

	p.Add(barsA)
	//p.Legend.Add("Group A", barsA)
	p.NominalX(keys...)

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "lengths.png"); err != nil {
		panic(err)
	}
}
