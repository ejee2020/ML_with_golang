package main 
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"io"
)

func main() { 
	f, err := os.Open("iris.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 5
	var rawCSVData [][]string 
	for { 
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil{
			log.Println(err)
			continue
		}
		rawCSVData = append(rawCSVData, record)
	}
	fmt.Println(rawCSVData)
}