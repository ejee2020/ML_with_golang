package main

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/kniren/gota/dataframe"
)

func main() { 
	irisFile, err := os.Open("iris.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	for _, colName := range irisDF.Names() { 
		if colName != "species" {
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float(){
				v[i] = floatVal
			}
		

			p, err := plot.New()
			if err != nil {
				log.Fatal(err)
			}
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			h.Normalize(1)
			p.Add(h)
			if err:= p.Save(4 *vg.Inch, 4 *vg.Inch, colName+"_hist.png"); err!= nil{
				log.Fatal(err)
		 }
		}
	}
}

