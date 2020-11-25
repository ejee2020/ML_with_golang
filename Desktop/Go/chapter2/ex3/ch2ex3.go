package main 
import (
	"fmt"
	"github.com/gonum/floats"
)
func main() { 
	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}
	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("A and B's dot product : %0.2f\n", dotProduct)

	floats.Scale(1.5, vectorA)
	fmt.Printf("A times 1.5: %v\n", vectorA)

	normB := floats.Norm(vectorB, 2)
	fmt.Printf("Vector B's Nomr: %0.2f\n", normB)
}