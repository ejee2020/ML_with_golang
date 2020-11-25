package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() { 
	f, err := os.Open("diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	diabetesDF := dataframe.ReadCSV(f)

	trainingNum := (4 * diabetesDF.Nrow()) / 5
	testNum := diabetesDF.Nrow() / 5
	if trainingNum + testNum < diabetesDF.Nrow() {
		trainingNum++
	}

	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i 
	}

	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	trainingDF := diabetesDF.Subset(trainingIdx)
	testDF := diabetesDF.Subset(testIdx)

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