package main 
import (

	"log"
	"os"
	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"
	"github.com/gonum/floats"
	"fmt"
	

)
func main() { 
	irisFile, err := os.Open("iris.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	sepalLength := irisDF.Col("sepal.length").Float() 

	minVal := floats.Min(sepalLength)
	maxVal := floats.Max(sepalLength)
	rangeVal := maxVal - minVal
	varianceVal := stat.Variance(sepalLength, nil)
	stdDevVal := stat.StdDev(sepalLength, nil)

	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)
	
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Max value: %0.2f\n", maxVal)
	fmt.Printf("Min value: %0.2f\n", minVal)
	fmt.Printf("Range value: %0.2f\n", rangeVal)
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n", stdDevVal)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	fmt.Printf("75 Quantile: %0.2f\n\n", quant75)

}
