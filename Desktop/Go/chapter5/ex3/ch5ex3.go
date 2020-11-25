package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	f, err = os.Create("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)

	for idx, record := range rawCSVData{
		if idx == 0 {
			if err := w.Write(record); err != nil {
				log.Fatal(err)
			}
			continue
		}

		outRecord := make([]string, 2)

		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		outRecord[0] = strconv.FormatFloat((score - 640.0)/(830.0 - 640.0), 'f', 4, 64)

		rate, err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		if err != nil {
			log.Fatal(err)
		}

		if rate <= 12.0 {
			outRecord[1] = "1.0"
			if err := w.Write(outRecord); err != nil {
				log.Fatal(err)
			}
			continue
		}
		outRecord[1] = "0.0"
		if err := w.Write(outRecord); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	}
	