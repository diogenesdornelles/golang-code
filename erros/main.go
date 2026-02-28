package main

import (
	"errors"
	"fmt"
)

/*
Não existem blocos try/catch

Em linguagens como Python ou Java, erros são Exceções (eventos que interrompem o fluxo e "sobem" na pilha até alguém capturá-los). Em Go, erros são Valores.
Essa é uma decisão de design profunda: o Go te obriga a encarar o erro como parte do fluxo lógico, não como algo que você "tenta" (try) e "pega" (catch) depois.

*/

// O tipo error é uma interface pré-definida em Go, que tem um método Error() string. Qualquer tipo que implemente esse método é considerado um erro.
type error interface {
	Error() string
}

// Isso significa que qualquer struct que tenha um método Error() retornando uma string é um erro válido.

// Criando um erro simples rapidamente, com retorno em tupla (valor, erro)
func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		// Criando um erro simples rapidamente
		return 0, errors.New("não é possível dividir por zero")
	}
	return a / b, nil // nil significa "sem erro"
}

type ErroConexao struct {
	Cod    int
	Server string
}

func (e ErroConexao) Error() string {
	return fmt.Sprintf("Erro %d no servidor %s", e.Cod, e.Server)
}

var ErrNotFound = errors.New("não encontrado") // Sentinel Error

func BuscarNoBanco() error {
	return fmt.Errorf("falha na query: %w", ErrNotFound) // Faz o wrapping
}

/*
Panic e Recover (O "Último Recurso")
Go tem o panic, que se assemelha a uma exceção, mas ele deve ser usado apenas para erros irrecuperáveis (ex: faltou memória, arquivo essencial corrompido).

panic: Para o programa imediatamente.

defer: Uma função que será executada ao final do escopo, aconteça o que acontecer.

recover: Captura um panic e impede o programa de fechar.


*/

func proteger() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de um desastre:", r)
		}
	}()
	panic("EXPLODIU TUDO!")
}

func main() {
	res, err := Dividir(10, 0)
	if err != nil {
		// Tratamento de erro (Logging, Retry, Return)
		fmt.Println("Erro detectado:", err)
		return
	}
	fmt.Println("Resultado:", res)

	_err := BuscarNoBanco()

	// Verificando se o erro original é o ErrNotFound
	if errors.Is(_err, ErrNotFound) {
		fmt.Println("Tratamento específico para 404")
	}

}
