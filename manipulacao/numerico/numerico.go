package numerico

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// String para Inteiro (Atoi: ASCII to Integer)
	numero, err := strconv.Atoi("123")

	if err != nil {
		fmt.Println("Erro ao converter string para inteiro:", err)
	} else {
		fmt.Println("Número convertido:", numero) // 123
	}

	// Float para String
	texto := strconv.FormatFloat(3.1415, 'f', 2, 64) // "3.14"
	fmt.Println("Float para String:", texto)
	// Matemática de alto nível
	fmt.Println(math.Sqrt(16))    // 4
	fmt.Println(math.Max(10, 20)) // 20
}
