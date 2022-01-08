package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		exibeIntro()
		//switch no go nao precisa de break
		comando := leComando()
		switch comando {
		case 1:
			fmt.Println("Iniciando monitoramento...")
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Encerrando...")
			os.Exit(0)
		default:
			fmt.Println("Comando nao reconhecido! ")
			os.Exit(-1)
		}

	}
}

func exibeIntro() {
	versao := 1.1
	fmt.Println("Software na versao", versao)
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir os logs")
	fmt.Println("0- Encerrar")
}
func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Endereco da variavel comando -> ", &comando) //memoria
	fmt.Println("O comando escolhido foi", comando)
	return comando
}
func iniciarMonitoramento() {
	site := "https://github.com"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, " foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, " nao foi carregado!")
		fmt.Println("Status code -> ", resp.StatusCode)
	}

}
