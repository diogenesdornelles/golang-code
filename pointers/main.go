// file: main.go

package main

import (
	"fmt"
)

type Person struct{ Name string }

func updateName(p *Person) {
	p.Name = "Novo Nome"
}

func f() *int {
	v := 1
	return &v
}

func incrementByRef(num *int) {
	*num++
}

func main() {

	// Ponteiros o básico: Permite acessar e modificar o valor de uma variável através de seu endereço de memória.
	x := 10
	p := &x                                              // p, of type *int, points to x
	fmt.Println("valor de x:", x)                        // "10"
	fmt.Println("Endereço de x através de &x em p:", p)  // "10"
	fmt.Println("valor de x através de *p:", *p)         // "10"
	*p = 20                                              // equivalent to x = 2
	fmt.Println("valor de x após modificação de *p:", x) // "20"

	// Ponteiros para Structs: Permite modificar campos de structs sem copiar a struct inteira.
	person := Person{Name: "Antigo"}
	updateName(&person)
	fmt.Println(person.Name) // Saída: Novo Nome

	n := 5
	incrementByRef(&n)
	fmt.Println(n) // Saída: 6

	// Testando zero value de ponteiros: O zero value de um ponteiro é nil, indicando que ele não aponta para nenhum endereço válido.
	var z, y int
	fmt.Println(&z == &z, &z == &y, &z == nil) // "true false false"
	// Each cal l of f returnsadistinc t value:
	fmt.Println(f() == f()) // "false"
}
