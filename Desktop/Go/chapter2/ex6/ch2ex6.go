package main 
import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	
)
func main() { 
	// Create a flat representation of our matrix.
	data := []float64{1.2, -5.7, -2.4, 7.3}

	// Form our matrix.
	a := mat64.NewDense(2, 2, data)

	val := a.At(0, 1)
	fmt.Printf("The value of ta at (0, 1) is: %.2f\n\n", val)

	col := mat64.Col(nil, 0, a)
	fmt.Printf("The values in the 1st column are: %v\n\n", col)

	row := mat64.Row(nil, 1, a)
	fmt.Printf("The values in the 2nd row are: %v\n\n", row)

	a.Set(0, 1, 11.2)
	
	a.SetRow(0, []float64{14.3, -4.2})

	a.SetCol(0, []float64{1.7, -0.3})
 
	fa := mat64.Formatted(a, mat64.Prefix("    "))
	fmt.Printf("A = %v\n\n", fa)
}