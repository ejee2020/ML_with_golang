package main

import (
	"log"
	"os"
	"github.com/kniren/gota/dataframe"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	
)

func main() { 
	advertFile, err := os.Open("Advertising.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer advertFile.Close()

	advertDF := dataframe.ReadCSV(advertFile)
	advertSummary := advertDF.Describe()
	fmt.Println(advertSummary)

	for _, colName := range advertDF.Names() {
		plotVals := make(plotter.Values, advertDF.Nrow())
		for i, floatVal := range advertDF.Col(colName).Float() { 
			plotVals[i] = floatVal
		}

		p, err := plot.New()
		if err != nil { 
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)
		
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		h.Normalize(1)
		p.Add(h)

		if err := p.Save(4 * vg.Inch, 4 * vg.Inch, colName+"_hist.png"); err != nil{
			log.Fatal(err)
		}
	}
}