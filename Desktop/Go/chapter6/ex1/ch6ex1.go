package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"github.com/gonum/floats"
)

type centroid []float64

func main() {

	// Pull in the CSV file.
	irisFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Define the names of the three separate species contained in the CSV file.
	speciesNames := []string{
		"Setosa",
		"Versicolor",
		"Virginica",
	}

	centroids := make(map[string]centroid)
	// Create a map to hold the filtered dataframe for each cluster.
	clusters := make(map[string]dataframe.DataFrame)

	
	for _, variety := range speciesNames {

		filter := dataframe.F{
			Colname: "variety",
			Comparator: "==",
			Comparando: variety,
		}

		filtered := irisDF.Filter(filter)
		clusters[variety] = filtered
		summaryDF := filtered.Describe()
		
		var c centroid

		for _, feature := range summaryDF.Names() {
			if feature == "column" || feature == "variety" {
				continue
			}
			c = append(c, summaryDF.Col(feature).Float()[0])
		}

		centroids[variety] = c
	}
	
	labels := irisDF.Col("variety").Records()
	floatColumns := []string{
		"sepal.length",
		"sepal.width",
		"petal.length",
		"petal.width",
	}

	var silhouette float64 
	for idx, label := range labels {
		var a float64 
		for i := 0; i < clusters[label].Nrow(); i++ {
			current := dfFloatRow(irisDF, floatColumns, idx)
			other := dfFloatRow(clusters[label], floatColumns, i)

			// Add to a.
			a += floats.Distance(current, other, 2) / float64(clusters[label].Nrow())
		}
		var otherCluster string
		var distanceToCluster float64
		
		for _, variety := range speciesNames{
			if variety == label {
				continue
			}

			distanceForThisCluster := floats.Distance(centroids[label], centroids[variety], 2)
            if distanceToCluster == 0.0 || distanceForThisCluster < distanceToCluster {
				otherCluster = variety
				distanceToCluster = distanceForThisCluster
			}
		}

		var b float64

		for i := 0; i < clusters[otherCluster].Nrow(); i++{
			current := dfFloatRow(irisDF, floatColumns, idx)
			other := dfFloatRow(clusters[otherCluster], floatColumns, i)
			b += floats.Distance(current, other, 2) / float64(clusters[otherCluster].Nrow())	
		}

		if a > b {
			silhouette += ((b - a) / a) / float64 (len(labels))
		}
		silhouette += ((b - a) / b) / float64(len(labels))
	}
	fmt.Printf("\nAverage Silhouette Coefficient: %0.2f\n\n", silhouette)
}

// dfFloatRow retrieves a slice of float values from a DataFrame
// at the given index and for the given column names.
func dfFloatRow(df dataframe.DataFrame, names []string, idx int) []float64 {
	var row []float64
	for _, name := range names {
		row = append(row, df.Col(name).Float()[idx])
	}
	return row
}
