package colecoes

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// Slices
	numeros := []int{10, 5, 8, 20}
	slices.Sort(numeros)             // Ordena o próprio slice (in-place)
	fmt.Println(slices.Max(numeros)) // 20

	// Maps
	m1 := map[string]int{"A": 1}
	m2 := map[string]int{"A": 1}
	fmt.Println(maps.Equal(m1, m2)) // true
}
