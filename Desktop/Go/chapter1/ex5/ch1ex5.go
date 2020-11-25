package main 
import (
	
	"fmt"
	"log"
	"os"
	
	"github.com/kniren/gota/dataframe"
)

func main() { 
	irisFile, err := os.Open("iris.csv")
	if err != nil { 
		log.Fatal(err)
	}
	defer irisFile.Close()
	irisDF := dataframe.ReadCSV(irisFile)
	fmt.Println(irisDF)
}