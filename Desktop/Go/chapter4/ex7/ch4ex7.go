package main

import (

	"log"
	"os"
	"encoding/csv"
	"fmt"
	"strconv"
	"github.com/gonum/matrix/mat64"
)

func main() { 
	f, err := os.Open("training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4
	
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	featureData := make([]float64, 4 * len(rawCSVData))
	yData := make([]float64, len(rawCSVData))

	var featureIndex int
	var yIndex int 

	for idx, record := range rawCSVData {
		if idx == 0 {
			continue
		}

		for i , val := range record {
			valParsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}

			if i < 3 {
				if i == 0 {
					featureData[featureIndex] = 1
					featureIndex++
				}

				featureData[featureIndex] = valParsed
				featureIndex++
			}
			if i == 3 {
				yData[yIndex] = valParsed
				yIndex++
			}
		}
	}

	features := mat64.NewDense(len(rawCSVData), 4, featureData)
	y := mat64.NewVector(len(rawCSVData), yData)

	if features != nil && y != nil {
		fmt.Println("Matrices formed for ridge regression")
	}

}