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
	loanDataFile, err := os.Open("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer loanDataFile.Close()

	// Create a dataframe from the CSV file.
	loanDF := dataframe.ReadCSV(loanDataFile)
	loanSummary := loanDF.Describe()
	fmt.Println(loanSummary)

	for _, colName:= range loanDF.Names() {
		plotVals := make(plotter.Values, loanDF.Nrow())
		for i , floatval := range loanDF.Col(colName).Float() {
			plotVals[i] = floatval
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

		if err := p.Save(4 * vg.Inch, 4 * vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}