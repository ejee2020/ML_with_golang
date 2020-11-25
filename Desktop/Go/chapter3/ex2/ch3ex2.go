package main
import (
    "os"
	"log"
	"encoding/csv"
	"io"
	"fmt"
	"strconv"
)

func main() { 
	f, err := os.Open("labeled.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	var observed []int
	var predicted []int

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

		observedVal, err := strconv.Atoi(record[0])
		if err !=  nil { 
			log.Printf("unexpected type in line ")
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil { 
			log.Printf("unexpected type")
			continue
		}

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	var truePosNeg int 
	for idx, oVal := range observed {
		if oVal == predicted[idx]{
			truePosNeg++
		}
	}

	accuracy := float64(truePosNeg) / float64(len(observed))

	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)

	classes := []int{0, 1, 2}

	for _, class := range classes {
		var truePos int
		var falsePos int 
		var falseNeg int 

		for idx, oval := range observed {
			switch oval {
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				}
				falseNeg++
		

		    default:
			    if predicted[idx] == class { 
				    falsePos++
			}

		}
	}
	precision := float64(truePos) / float64(truePos + falsePos)
	recall := float64(truePos) / float64(truePos + falseNeg)
	fmt.Printf("\n Accuracy (Class %d) = %0.2f", class, precision)
	fmt.Printf("\n recall (Class %d) = %0.2f\n\n", class, recall)
}
}