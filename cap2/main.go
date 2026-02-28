package main

import (
	"fmt"
	"github.com/diogenesdornelles/cap-2-lessons/finance"
)

// Package-level declaration: Visível em todo o pacote 'main'
var versao = "1.26.0" 

func main() {
    // 1. Short Variable Declaration (:=)
    // Usado para inicializar variáveis locais. O tipo é inferido.
    ingressoBase := finance.Dinheiro(10000) // R$ 100,00

    // 2. Assignments (=) e Reatribuição
    // 'valorFinal' nasce aqui.
    valorFinal := ingressoBase + finance.TaxaDeServico

    // 3. Múltiplas atribuições (Tuplas) - estilo Go/Python
    // Útil para swap de valores sem variável temporária (como no Assembly!)
    a, b := 1, 2
    a, b = b, a // Swap atômico no nível da linguagem

    // 4. Pointers (Ponteiros)
    // Assim como o endereço no registrador, o ponteiro guarda a localização.
    p := &valorFinal 
    fmt.Println("Endereço na RAM:", p)
    *p = *p + 200 // Modificando o valor original via desreferenciação

    // 5. Scope (Escopo)
    fmt.Printf("Sistema v%s | Preço: %s\n", versao, valorFinal)

    {
        // Escopo de bloco: Esta variável 'temp' morre no '}'
        temp := "Promoção Relâmpago"
        fmt.Println(temp)
    }
    // fmt.Println(temp) // Erro de compilação! 'temp' não existe aqui.

    demonstrarShadowing()
}

func demonstrarShadowing() {
    x := "Original"
    if true {
        x := "Sombreado" // CUIDADO: ':=' aqui cria uma NOVA variável no escopo do IF
        fmt.Println("Dentro do IF:", x)
    }
    fmt.Println("Fora do IF:", x) // Imprime "Original"
}