package main

import (
	"fmt"
	"log"
	"math"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	irisData, err := base.ParseCSVToInstances("iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	knn := knn.NewKnnClassifier("euclidean", "linear", 2)

	// Use cross-fold validation to successively train and evalute the model
	// on 5 folds of the data set.
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, knn, 5)
	if err != nil {
		log.Fatal(err)
	}
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	fmt.Printf("\nAccuracy\n%.2f (+/- %.2f)\n\n", mean, stdev*2)
}