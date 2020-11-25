package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	driverDataFile, err := os.Open("fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer driverDataFile.Close()
	
	driverDF := dataframe.ReadCSV(driverDataFile)

	driverSummary := driverDF.Describe()
	fmt.Println(driverSummary)

	for _, colName := range driverDF.Names() {
		plotVals := make(plotter.Values, driverDF.Nrow())
		for i , floatVal := range driverDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histrogram of %s", colName)
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		h.Normalize(1)
		p.Add(h)
		if err := p.Save(4 * vg.Inch, 4 * vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}