package texto

import (
	"fmt"
	"strings"
)

func main() {
	s := "  Golang é Incrível  "

	fmt.Println(strings.TrimSpace(s))      // "Golang é Incrível"
	fmt.Println(strings.ToUpper(s))        // "  GOLANG É INCRÍVEL  "
	fmt.Println(strings.Contains(s, "Go")) // true

	frutas := []string{"Maçã", "Uva", "Banana"}
	fmt.Println(strings.Join(frutas, "-")) // "Maçã-Uva-Banana"
}
