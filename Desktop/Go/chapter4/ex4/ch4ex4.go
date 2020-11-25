package main

import (
	"log"
	"os"
	"fmt"
	"encoding/csv"
	"strconv"
	"math"
	"github.com/sajari/regression"

)

func main() { 
	f, err := os.Open("training.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = 4 
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var r regression.Regression 
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
   
	for i , record := range trainingData { 
		if i == 0 {
			continue
		}

		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		r.Train(regression.DataPoint(yVal, []float64{tvVal}))
	}

	r.Run()
	fmt.Printf("\nRegression Formula : \n%v\n\n", r.Formula)

	f, err = os.Open("test.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer f.Close()

	reader = csv.NewReader(f)

	reader.FieldsPerRecord = 4 
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var mAE float64
	for i , record := range testData{
		if i == 0 {
			continue
		}

		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		yPredicted, err := r.Predict([]float64{tvVal})

		mAE += math.Abs(yObserved - yPredicted) / float64(len(testData))
	}
	fmt.Printf("MAE = %0.2f\n\n", mAE)
}