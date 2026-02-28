package data

/* 
Tipos básicos em Go

Declaração fora de função, no nível do pacote (global), usa-se = para atribuição

Em nível de função, usa-se := para declaração e atribuição simultânea (inferência de tipo)

O Conceito de "Zero Value"
No Go, se você não atribui valor, o compilador limpa a memória para você (diferente do C/Assembly onde você pode ler lixo).

Números: 0

Booleanos: false

Strings: ""

Ponteiros, Slices, Maps, Channels, Interfaces: nil

*/

// Byte: Alias para uint8. Usado para dados brutos (Raw Data)
var Byte byte = 255

// inteiro com o padrão do tipo int (32 ou 64 bits dependendo da arquitetura)
var Inteiro = 42

var Inteiro64 int64 = 1234567890123456789
var Inteiro32 int32 = 1234567890
var Inteiro16 int16 = 12345
var Inteiro8 int8 = 123

var Inteiro64u uint64 = 1234567890123456789
var Inteiro32u uint32 = 1234567890
var Inteiro16u uint16 = 12345
var Inteiro8u uint8 = 123


// O padrão é float64, mas também existe float32
var Flutuante = 3.14
var Flutuante64 float64 = 3.141592653589793
var Flutuante32 float32 = 3.14

// Complexos: Usados para computação científica (parte real e imaginária)
var Complexo64 complex64 = 1 + 2i
var Complexo128 complex128 = 3.14 + 2.71i

// uintptr: Um tipo inteiro grande o suficiente para guardar um endereço de memória.
// Muito usado em operações "Unsafe" para aritmética de ponteiros (estilo C).
var EnderecoMemoria uintptr = 0xDEADBEEF

var Booleano = true
var Caractere = 'A'

var Rune rune = 'R' // rune é um alias para int32, usado para representar caracteres Unicode
var String = "Olá, Go!"

// Arrays e Slices

// Array com tamanho definido pelo compilador (...)
var ArrayAuto = [...]int{10, 20, 30, 40}

var Array [3]int = [3]int{1, 2, 3} // Array de tamanho fixo

// Array Multidimensional (Matriz)
// Na memória: um bloco contínuo de 4 * 8 bytes (se int64)
var ArrayMatriz [2][2]int = [2][2]int{
    {1, 2},
    {3, 4},
}

var SliceDin []int = []int{4, 5, 6}    // Slice de tamanho dinâmico

var SliceLiteral []string = []string{"CPU", "RAM", "SSD"}

var SliceMake []string = make([]string, 0)

// Nil Slice vs Empty Slice
var SliceNil []int          // pointer=nil, len=0, cap=0 (não alocou nada)
var SliceVazio = []int{}    // pointer=0x..., len=0, cap=0 (alocou um cabeçalho vazio)

// Variando o Make
// make([]T, len, cap) permite criar um slice com capacidade maior que o comprimento inicial, o que pode melhorar a performance ao adicionar elementos.
// Aqui criamos um slice de inteiros com comprimento 5 e capacidade 10. Ele tem espaço para crescer sem precisar realocar imediatamente.
var SliceComBuffer = make([]int, 5, 10)


// Mapas
var CountMake map[string]int = make(map[string]int)

var CountLiteral = map[string]int{"CPU": 10, "RAM": 20, "SSD": 5}

// Inicialização segura
// Em Go, mapas não inicializados (nil) causam panic se você tentar adicionar chaves. Sempre use make ou literal para evitar isso.
var Precos = make(map[string]int, 100) // Pré-aloca espaço para 100 chaves (Performance!)

// Definição de struct simples
type Person struct {
    Name string
    Age  int
}

var Pessoa = Person{Name: "Alice", Age: 30}


// Interface Vazia: Pode conter QUALQUER valor.
var QualquerCoisa interface{} = "Pode ser string, int, ou struct"

// No Go 1.18+, 'any' é um alias oficial para 'interface{}'
var QualquerCoisaModerna any = 42

// Em Go, funções são tipos. Você pode guardá-las em variáveis.
var Operacao func(int, int) int = func(a, b int) int {
    return a + b
}

// Chan: Usado para comunicação entre threads (Goroutines).
// É como um pipe sincronizado no nível do hardware.
var CanalSincronizado chan int = make(chan int)

// Error: Um tipo de interface embutido para tratamento de erros.
var ErroDeSistema error = nil 

// Pointer: Ponteiro para outro tipo.
var PonteiroParaInt *int = nil // Inicialmente nil, pode apontar para um inteiro mais tarde


// Mode represents the application mode (dev or prod)
type Mode string

// Constants for Mode enum
const (
    Dev  Mode = "dev"
    Prod Mode = "prod"
)

// Config holds all application configurations
type Config struct {
    Mode    Mode
    // Add other configs here, e.g., DatabaseURL, StripeKey, etc.
}

var AppConfig = Config{
    Mode: Dev, // Ou Prod para produção
}