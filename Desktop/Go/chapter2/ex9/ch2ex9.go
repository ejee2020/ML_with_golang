package main 
import (
	"fmt"
	"log"
	"os"
	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"
	"github.com/montanaflynn/stats"

)
func main() { 
	irisFile, err := os.Open("iris.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	sepalLength := irisDF.Col("petal.length").Float() 

	meanVal := stat.Mean(sepalLength, nil)

	modeVal, modeCount := stat.Mode(sepalLength, nil)

	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n Sepal Length stats: \n")
	fmt.Printf("Mean: %0.2f\n", meanVal)
	fmt.Printf("Mode: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	fmt.Printf("Median: %0.2f\n\n", medianVal)
}


