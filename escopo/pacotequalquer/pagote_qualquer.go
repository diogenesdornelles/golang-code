package pacotequalquer


func Signum(x int) int {
	// A constante Pi do pacote main não é acessível aqui, pois main não é um pacote importado, mas sim o pacote principal do programa. Cada pacote tem seu próprio escopo, e as constantes, variáveis e funções definidas em um pacote não são acessíveis diretamente em outro pacote, a menos que sejam exportadas (começando com letra maiúscula) e o pacote seja importado.
	//fmt.Println("Pi from main:", main.Pi) // Acessando a constante Pi do pacote main
	
	switch {
	case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
	}
}
