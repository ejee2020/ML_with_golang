package main

import (
	"log"
	"os"
	"github.com/kniren/gota/dataframe"
	"bufio"
)

func main() { 
	f, err := os.Open("Advertising.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer f.Close()

	advertDF := dataframe.ReadCSV(f)
	trainingNum :=  (4 * advertDF.Nrow()) / 5
	testNum := advertDF.Nrow() / 5
	if trainingNum + testNum < advertDF.Nrow() {
		trainingNum++
	}

	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	for i :=0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	trainingDF := advertDF.Subset(trainingIdx)
	testDF := advertDF.Subset(testIdx)

	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	for idx, setName := range []string{"training.csv", "test.csv"} {
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		w := bufio.NewWriter(f)
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}


