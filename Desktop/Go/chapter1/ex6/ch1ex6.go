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
	
	filter := dataframe.F{
		Colname: "variety",
		Comparator: "==",
		Comparando: "Versicolor",
	}

	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}
	fmt.Println(versicolorDF)

	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal.width", "variety"})
	fmt.Println(versicolorDF)

	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal.width", "variety"}).Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)

}