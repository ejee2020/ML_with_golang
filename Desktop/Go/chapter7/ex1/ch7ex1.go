package main

import (
    "image/color"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	passengersFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()
	
	passengersDF := dataframe.ReadCSV(passengersFile)
	yVals := passengersDF.Col("AirPassengers").Float()
	pts := make(plotter.XYs, passengersDF.Nrow())

	for i , floatVal := range passengersDF.Col("time").Float(){
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "passengers"
	p.Add(plotter.NewGrid())

	l, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p.Add(l)
	if err := p.Save(10 * vg.Inch, 4 * vg.Inch, "passengers_ts.png"); err != nil {
		log.Fatal(err)
	}
}