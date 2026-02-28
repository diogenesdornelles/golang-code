package chars

import (
	"fmt"
	"unicode"
)

func main() {
	r := '7'

	fmt.Println(unicode.IsDigit(r))   // true
	fmt.Println(unicode.IsLetter(r))  // false
	fmt.Println(unicode.ToUpper('ç')) // 'Ç'
}
