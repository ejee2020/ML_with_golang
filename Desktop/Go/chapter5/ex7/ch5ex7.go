package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	
	"os"
	"strconv"
	"io"
)

func main(){
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	var observed []float64
	var predicted []float64

	line := 1

	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Skip the header.
		if line == 1 {
			line++
			continue
		}

		// Read in the observed value.
		observedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Make the corresponding prediction.
		score, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal := predict(score)

		// Append the record to our slice, if it has the expected type.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	var truePosNeg int

	// Accumulate the true positive/negative count.
	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}

	// Calculate the accuracy (subset accuracy).
	accuracy := float64(truePosNeg) / float64(len(observed))

	// Output the Accuracy value to standard out.
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}

// predict makes a prediction based on our
// trained logistic regression model.
func predict(score float64) float64 {

	// Calculate the predicted probability.
	p := 1 / (1 + math.Exp(-13.65*score+4.89))

	// Output the corresponding class.
	if p >= 0.5 {
		return 1.0
	}

	return 0.0
}