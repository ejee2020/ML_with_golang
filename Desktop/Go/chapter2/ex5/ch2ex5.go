package main 
import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	
)
func main() { 
	data := []float64{1.2, -5.7, -2.4, 7.3}
	a := mat64.NewDense(2, 2, data)
	fa := mat64.Formatted(a, mat64.Prefix(" "))
	fmt.Printf("mat = %v\n\n", fa)
}