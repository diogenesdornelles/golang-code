// file: main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Cria um mapa para armazenar cada linha como chave e o número de ocorrências como valor.
	counts := make(map[string]int)
	// Cria um scanner para ler da entrada padrão
	input := bufio.NewScanner(os.Stdin)
	//  Loop que lê linhas uma por uma até o fim da entrada (EOF). Cada linha é adicionada ou incrementada no mapa.
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	// EOF Ctrl + D (Linux/Mac) ou Ctrl + Z (Windows)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
