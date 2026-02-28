package forloop

import "fmt"


// Exportação em maiúsculo torna a função acessível fora do pacote.
func Run() {
    // 1. O "For Tradicional" (Estilo C / Assembly)
    // Tem: inicialização; condição; pós-execução
    fmt.Println("--- For Clássico ---")
    for i := 0; i < 3; i++ {
        fmt.Printf("Contagem: %d\n", i)
    }

    // 2. O "For-While" (Apenas a condição)
    // Em Go, se você omitir o início e o fim, o 'for' vira um 'while'.
    fmt.Println("\n--- For estilo While ---")
    count := 0
    for count < 3 { 
        fmt.Printf("While: %d\n", count)
        count++
    }

    // 3. O "For Infinito" (O famoso loop JMP do Assembly)
    // Equivalente ao 'while(true)' ou 'for(;;)'
    fmt.Println("\n--- For Infinito com Break ---")
    i := 0
    for { 
        if i >= 3 {
            break // Sai do loop (como um JZ/JNZ no hardware)
        }
        fmt.Printf("Looping... %d\n", i)
        i++
    }

    // 4. O "For Range" (Iterador)
    // Usado para percorrer Slices, Maps, Strings ou Channels.
    // É o mais seguro e performático para coleções.
    fmt.Println("\n--- For Range (Slice) ---")
    items := []string{"CPU", "RAM", "SSD"}
    for index, value := range items {
        fmt.Printf("Índice [%d] tem o hardware: %s\n", index, value)
    }

    // 4.1 For Range (Apenas valor, ignorando o índice)
    for _, v := range items {
        fmt.Println("Hardware:", v)
    }
}