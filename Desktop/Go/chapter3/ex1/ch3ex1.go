package main
import (
    "os"
	"log"
	"encoding/csv"
	"io"
	"fmt"
	"math"
	"strconv"
	"gonum.org/v1/gonum/stat"
)

func main() { 
	f, err := os.Open("continuous_data.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer f.Close() 

	reader := csv.NewReader(f)

	var observed []float64 
	var predicted []float64 

	line := 1

	for { 
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if line == 1 {
			line++
			continue
		}
		
		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil { 
			log.Printf("Unexpected type in line number %d\n", line)
			continue
		}

		predictedVal , err := strconv.ParseFloat(record[1], 64)
		if err != nil { 
			log.Printf("Unexpected type in line number %d\n", line)
			continue
		}

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line ++
	}

	var mAE float64
	var mSE float64 
	for idx, oVal := range observed { 
		mAE += math.Abs(oVal - predicted[idx]) / float64(len(observed))
		mSE += math.Pow(oVal - predicted[idx], 2) / float64(len(observed))
	}

	fmt.Printf("\nMAE = %0.2f\n", mAE)
	fmt.Printf("\nMSE = %0.2f\n\n", mSE)

	rSquared := stat.RSquaredFrom(observed, predicted, nil)
	fmt.Printf("\nR^2 = %0.2f\n\n", rSquared)
}