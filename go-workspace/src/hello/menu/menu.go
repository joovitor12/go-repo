package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 5
const delay = 5

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
	fmt.Println("")
	return comando
}
func iniciarMonitoramento() {
	sites := []string{"https://github.com", "https://www.linkedin.com", "https://mail.google.com/mail/u/0/#inbox", "https://www.youtube.com"}
	//vezes que o loop sera repetido
	for i := 0; i < monitoramentos; i++ {
		for i, j := range sites {
			fmt.Println(i, j)
			testandoSite(j)
			fmt.Println("")
		}
		//delay
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")

}

func testandoSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, " foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, " nao foi carregado!")
		fmt.Println("Status code -> ", resp.StatusCode)
	}

}
