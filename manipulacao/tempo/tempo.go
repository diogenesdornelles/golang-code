package tempo

import (
	"fmt"
	"time"
)

// Agendar demonstra o uso de time.Sleep para pausas síncronas.
// Útil para simular delays em execuções sequenciais.
func Agendar() {
	time.Sleep(2 * time.Second)
	fmt.Println("2 segundos passaram!")
}

// ExibirTempo demonstra várias funcionalidades básicas de manipulação de tempo em Go,
// incluindo formatação, timestamps, comparações, durações e conversões de fuso.
func ExibirTempo() {
	// Formatação de data
	fmt.Println("Formatação de data pt-br:")
	agora := time.Now()
	fmt.Println(agora.Format("02/01/2006")) // Data formatada (DD/MM/YYYY)

	// Timestamp Unix
	fmt.Println("Timestamp Unix:")
	fmt.Println(agora.Unix())

	// Manipulação de datas: adicionar dias e comparações
	fmt.Println("Manipulação de datas: adicionar dias e comparações")
	amanha := agora.AddDate(0, 0, 1)
	fmt.Println(amanha.Before(agora)) // false (amanhã é depois de agora)

	ontem := agora.AddDate(0, 0, -1)
	fmt.Println(ontem.Format("02/01/2006")) // Data formatada

	// Durações: criação e exibição
	fmt.Println("Durações: criação e exibição")
	duracao := 2*time.Hour + 30*time.Minute
	fmt.Println(duracao.String()) // 2h30m0s

	// Criação de datas com fuso fixo (UTC-3)
	fmt.Println("Criação de datas com fuso fixo (UTC-3)")
	data1 := time.Date(2024, 6, 1, 0, 0, 0, 0, time.FixedZone("", -3*3600))
	fmt.Println(data1)            // Exibe no fuso local
	fmt.Println(data1.UTC())      // Exibe em UTC
	fmt.Println(data1.Location()) // Mostra o fuso (FixedZone)

	data2 := time.Date(2024, 6, 15, 0, 0, 0, 0, time.FixedZone("", -3*3600))
	diferenca := data2.Sub(data1)
	fmt.Printf("Diferença em dias: %.0f\n", diferenca.Hours()/24) // Diferença em dias

	// Conversão para outro fuso horário
	fmt.Println("Conversão para outro fuso horário: LoadLocation e In")
	loc, _ := time.LoadLocation("America/Sao_Paulo") // UTC-3
	original := time.Now()
	convertido := original.In(loc) // Converte para o novo fuso
	fmt.Println(convertido)

	// Parsing de string para time.Time
	fmt.Println("Parsing de string para time.Time")
	parsed, err := time.Parse("02/01/2006", "15/06/2024")
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Data parsed:", parsed)
	}

	// Comparações de tempos
	fmt.Println("Comparações de tempos")
	t1 := time.Now()
	t2 := t1.Add(time.Hour)
	fmt.Println(t2.After(t1)) // true
	fmt.Println(t1.Equal(t1)) // true

	// Manipulação de precisão: Round e Truncate
	fmt.Println("Manipulação de precisão: Round e Truncate")
	_agora := time.Now()
	arredondado := _agora.Round(time.Hour)   // Arredonda para a hora mais próxima
	truncado := _agora.Truncate(time.Minute) // Trunca para o minuto
	fmt.Println("Arredondado:", arredondado)
	fmt.Println("Truncado:", truncado)

	// Timestamps avançados
	fmt.Println("Timestamps avançados")
	fmt.Println("Unix Nano:", time.Now().UnixNano())

	// Fuso horário customizado
	fmt.Println("Fuso horário customizado")
	customZone := time.FixedZone("Custom", 5*3600) // UTC+5
	customTime := time.Date(2024, 6, 1, 12, 0, 0, 0, customZone)
	fmt.Println("Custom fuso:", customTime)
}
