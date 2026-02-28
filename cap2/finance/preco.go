package finance

import "fmt"

// Type Declaration: um tipo novo baseado em int64.
// Isso evita que some "Dinheiro" com "Idade" acidentalmente.
type Dinheiro int64 

// Names: Constantes exportadas (PascalCase)
const TaxaDeServico Dinheiro = 500 // 5.00 em centavos

// Método associado ao tipo (comportamento)
func (d Dinheiro) String() string {
    return fmt.Sprintf("R$ %.2f", float64(d)/100)
}

// Declaration: Uma variável de pacote (nível global do package)
var TotalVendas Dinheiro