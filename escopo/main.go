// file: main.go

package main

import (
	"fmt"
	"github.com/diogenesdornelles/escopo/pacotequalquer"
)

/*
Reforçando conceitos de escopo: Em Go, o escopo controla a visibilidade de identificadores (variáveis, constantes, funções, etc.). Há três níveis principais:

Escopo universal: Identificadores pré-definidos (built-ins) como int, fmt.Println, etc., visíveis em todo o programa.
Escopo de pacote: Identificadores declarados no nível do pacote (fora de funções) são visíveis em todo o pacote. Se começarem com maiúscula, são exportados e podem ser acessados por outros pacotes via importação.
Escopo de bloco: Identificadores declarados dentro de blocos (como if, for, funções) são visíveis apenas nesse bloco e sub-blocos. Eles "sombreiam" identificadores de escopos externos com o mesmo nome.

*/

// f() não pode ser visto em pacotequalquer
func f() int {
	return 0
}

func g(x int) int {
	return x + 1
}

// Declarar identificadores com letra maiúscula (exportados) no pacote main oferece pouco benefício prático, pois main não pode ser importado por outros pacotes
const (
	Pi = 3.14
)

// As funções não são visíveis fora do pacote, a menos que sejam exportadas (começando com letra maiúscula).	
// A constante Pi é visível em todo o pacote, mas não fora dele, pois começa com letra maiúscula, indicando que é exportada.

func main() {
	fmt.Println("Starting app.")

	res :=pacotequalquer.Signum(10)
	fmt.Println("Signum:", res)
	// Variáveis declaradas dentro de um bloco (como IF, ELSE, FOR) são visíveis apenas dentro desse bloco e seus blocos aninhados.
	// Escopo implícito do primeiro IF
	x := f()
	if x == 0 {
		fmt.Println(x)
	} else {
		{ // Escopo implícito do ELSE IF
			y := g(x)
			if x == y {
				fmt.Println(x, y)
			} else {
				fmt.Println(x, y) // y ainda é visível aqui!
			}
		}
	}
}
