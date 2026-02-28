// file: main.go

/*
1. Introdução ao fmt.Printf
fmt.Printf permite formatar strings com placeholders chamados "verbos".
Os verbos começam com % e indicam o tipo de valor a ser inserido.
Exemplos comuns (baseados no comentário no seu main.go):
%d: Inteiro decimal.
%x, %o, %b: Inteiro em hexadecimal, octal ou binário.
%f, %g, %e: Números flutuantes.
%t: Booleano (true/false).
%c: Caractere (rune).
%s: String.
%q: String ou rune entre aspas.
%v: Valor em formato natural (genérico).
%T: Tipo do valor.
%%: Sinal de porcentagem literal.
2. Exemplos com Variáveis do data.go
Usaremos as variáveis exportadas (maiúsculas) do pacote data. Aqui estão exemplos de impressão:

Inteiros: %d para decimal, %x para hex, etc.
Flutuantes: %f para ponto fixo, %e para notação científica.
Booleanos e Caracteres: %t para bool, %c para rune.
Strings: %s para string simples, %q para entre aspas.
Arrays/Slices: %v para imprimir o conteúdo.
Maps e Structs: %v para formato natural.
Outros: %T para tipo, %v para genérico.


3. Diferenças Práticas entre Métodos do Pacote fmt em Go
O pacote fmt oferece várias funções para entrada/saída formatada. Aqui estão as principais diferenças práticas, focando em uso comum:

Printf: Formata uma string com verbos (ex.: %s, %d) e imprime na saída padrão (stdout). Ideal para mensagens formatadas no console. Exemplo: fmt.Printf("Nome: %s, Idade: %d\n", nome, idade).

Println: Imprime valores separados por espaço, adicionando uma nova linha no final. Não usa formatação; apenas concatena. Exemplo: fmt.Println("Olá", nome, idade) (saída: "Olá Alice 30\n").

Fprintf: Similar ao Printf, mas escreve em um io.Writer (arquivo, buffer, rede, etc.), não na stdout. Útil para logs ou arquivos. Exemplo: fmt.Fprintf(file, "Erro: %v\n", err).

Outros Métodos Comuns:

Sprintf: Retorna uma string formatada em vez de imprimir. Exemplo: str := fmt.Sprintf("Total: %.2f", total).
Scanf: Lê entrada formatada da stdin. Exemplo: fmt.Scanf("%s %d", &nome, &idade).
Fscanf: Lê de um io.Reader formatado. Exemplo: fmt.Fscanf(file, "%s", &linha).
Print / Fprint / Sprint: Versões sem nova linha automática (Println adiciona).
*/


package main

import (
	"fmt"
   "github.com/diogenesdornelles/impressao/data"
   "os"
)

// ...existing code...

func main() {

   // Exemplos de impressão com fmt.Printf usando variáveis do pacote data

   // Inteiros
   fmt.Printf("Byte (uint8): %d\n", data.Byte)
   fmt.Printf("Inteiro: %d\n", data.Inteiro)
   fmt.Printf("Inteiro64: %d (em hex: %x)\n", data.Inteiro64, data.Inteiro64)
   fmt.Printf("Inteiro32u: %d (em octal: %o)\n", data.Inteiro32u, data.Inteiro32u)

   // Flutuantes
   fmt.Printf("Flutuante: %f\n", data.Flutuante)
   fmt.Printf("Flutuante64: %g (notação científica: %e)\n", data.Flutuante64, data.Flutuante64)

   // Complexos
   fmt.Printf("Complexo128: %v\n", data.Complexo128)

   // Booleanos e Caracteres
   fmt.Printf("Booleano: %t\n", data.Booleano)
   fmt.Printf("Caractere: %c\n", data.Caractere)
   fmt.Printf("Rune: %c (código: %d)\n", data.Rune, data.Rune)

   // Strings
   fmt.Printf("String: %s\n", data.String)
   fmt.Printf("String entre aspas: %q\n", data.String)

   // Arrays e Slices
   fmt.Printf("ArrayAuto: %v\n", data.ArrayAuto)
   fmt.Printf("SliceDin: %v\n", data.SliceDin)
   fmt.Printf("SliceLiteral: %v\n", data.SliceLiteral)

   // Maps
   fmt.Printf("CountLiteral: %v\n", data.CountLiteral)

   // Struct
   fmt.Printf("Pessoa: %v\n", data.Pessoa)

   // Interface e Any
   fmt.Printf("QualquerCoisa: %v (tipo: %T)\n", data.QualquerCoisa, data.QualquerCoisa)
   fmt.Printf("QualquerCoisaModerna: %v (tipo: %T)\n", data.QualquerCoisaModerna, data.QualquerCoisaModerna)

   // Enum

   fmt.Printf("appMode: %v\n", data.AppConfig.Mode)

   // Função
   fmt.Printf("Resultado da Operacao(5,3): %d\n", data.Operacao(5, 3))

   // Canal (não imprimível diretamente, mas podemos mostrar o tipo)
   fmt.Printf("CanalSincronizado: %T\n", data.CanalSincronizado)

   // Erro
   fmt.Printf("ErroDeSistema: %v\n", data.ErroDeSistema)

   // Ponteiro
   fmt.Printf("PonteiroParaInt: %v\n", data.PonteiroParaInt)

   // Demonstrações de outros métodos fmt (excluindo Printf)

   // Println: Imprime sem formatação, com nova linha
   fmt.Println("Println - Valores:", data.Inteiro, data.Flutuante, data.String)

   // Print: Imprime sem nova linha automática
   fmt.Print("Print - Sem nova linha: ", data.Booleano, " ")
   fmt.Println("Continuando...")

   // Fprintf: Escreve em um io.Writer (ex.: arquivo ou buffer)
   file, _ := os.Create("output.txt")
   defer file.Close()
   fmt.Fprintf(file, "Fprintf - Escrevendo no arquivo: %v\n", data.Pessoa)

   // Sprintf: Retorna uma string formatada
   str := fmt.Sprintf("Sprintf - String formatada: %s tem %d anos", data.Pessoa.Name, data.Pessoa.Age)
   fmt.Println(str)

   // Sprint: Retorna string sem formatação
   str2 := fmt.Sprint("Sprint - Sem formatação: ", data.ArrayAuto)
   fmt.Println(str2)

   // Fprint: Escreve em Writer sem nova linha automática
   fmt.Fprint(file, "Fprint - Sem nova linha: ", data.SliceDin)

   // Scanf: Lê entrada formatada da stdin (exemplo interativo)
   var nome string
   var idade int
   fmt.Print("Digite nome e idade (ex.: Alice 30): ")
   fmt.Scanf("%s %d", &nome, &idade)
   fmt.Printf("Lido: %s, %d\n", nome, idade)

   // Fscanf: Lê de um io.Reader (ex.: arquivo)
   file2, _ := os.Open("output.txt")
   defer file2.Close()
   var linha string
   fmt.Fscanf(file2, "%s", &linha)
   fmt.Println("Fscanf - Lido do arquivo:", linha)

   fmt.Printf("Hello, %s!\n", "world")


}
// ...existing code...