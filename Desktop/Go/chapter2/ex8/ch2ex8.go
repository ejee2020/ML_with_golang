package main 
import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"log"
	
)
func main() { 
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	
	ft := mat64.Formatted(a.T(), mat64.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	deta := mat64.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	aInverse := mat64.NewDense(0, 0, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}

	fi := mat64.Formatted(aInverse, mat64.Prefix("       "))
	fmt.Printf("a^-1 = %v\n\n", fi)
}