package main

import (

	"log"
	"os"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/kniren/gota/dataframe"
)

func main() { 
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	advertDF := dataframe.ReadCSV(f)

	yVals := advertDF.Col("Sales").Float()

	pts := make(plotter.XYs, advertDF.Nrow())

	ptsPred := make(plotter.XYs, advertDF.Nrow())

	for i, floatVal := range advertDF.Col("TV").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.X.Label.Text = "TV"
	p.Y.Label.Text = "Sales"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Radius = vg.Points(3)

	l, err := plotter.NewLine(ptsPred)
	if err != nil{
		log.Fatal(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	p.Add(s, l)
	if err := p.Save(4 * vg.Inch, 4 * vg.Inch, "regression_line.png"); err != nil {
		log.Fatal(err)
	}
}

func predict(tv float64) float64 {
	return 7.07 + tv * 0.05
}