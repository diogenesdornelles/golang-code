package colecoes

import (
	"fmt"
	"maps"
	"slices"
)

/*
Aqui está a anatomia de um Slice na memória. 

Ele é uma struct de 3 campos (24 bytes em sistemas 64-bit):

Pointer: Endereço do primeiro elemento do array que o slice "enxerga".

Len (Length): Quantos elementos o slice contém atualmente.

Cap (Capacity): Quantos elementos existem entre o início do slice e o fim do array subjacente.


O Array é o dado bruto na memória. O Slice é que possui essa estrutura de Header (Pointer, Len, Cap). 
Quando você passa um Array para uma função, o Go copia todos os elementos. 
Quando você passa um Slice, o Go copia apenas esses 24 bytes, tornando a passagem de Slices ultra eficiente.

a := [...]int{0, 1} -> O compilador cria um array fixo de 2 posições.

s := []int{0, 1} -> O compilador cria um array oculto (anônimo) de 2 posições e cria um Slice Header apontando para ele.

*/

func main() {
	
	// Arrays: tem tamanho fixo e tipo definido
	var arr [5]int
	arr[0] = 10
	arr[1] = 5
	arr[2] = 8
	arr[3] = 20
	arr[4] = 15

	fmt.Println(arr) // [10 5 8 20 15]
	preenchido := [...]int{99: -1}

	fmt.Println(preenchido) // [0 0 0 ... -1] preenchido[99] = -1, os outros elementos são 0
	
	a := [...]int{0, 1}
	s := [2]int{0, 1}

	fmt.Println(a == s) // true, arrays comparáveis se tiverem o mesmo tipo e elementos iguais
	
	// Slices: são arrays mutáveis e dinâmicos. 
	numeros := []int{10, 5, 8, 20}
	fmt.Println(len(numeros)) // 4
	fmt.Println(cap(numeros))
	slices.Sort(numeros)             // Ordena o próprio slice (in-place)
	fmt.Println(slices.Max(numeros)) // 20
	numeros2 := []int{10, 5, 8, 20}

	fmt.Println(slices.Equal(numeros, numeros2)) // true comparação de slices usando slices.Equal, que verifica se os elementos são iguais e na mesma ordem

	// Maps: são coleções de chave-valor

	m1 := map[string]int{"A": 1}
	m2 := map[string]int{"A": 1}
	fmt.Println(maps.Equal(m1, m2)) // true
}
