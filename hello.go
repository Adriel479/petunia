package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	monitoramentos = 3
	delay          = 5 * time.Second
)

func main() {
	exibirIntroducao()
	for {
		exibirMenu()
		comando := lerComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
		}
	}
}

func exibirIntroducao() {
	nome := "Boris"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibirMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{
		"https://alura.com.br",
		"https://google.com.br",
		"https://brisanet.com.br/teste",
		"https://youtube.com.br/abc",
	}
	for ii := 0; ii < monitoramentos; ii++ {
		for site := range sites {
			fmt.Println("Testando site", site, ":", sites[site])
			testarSite(sites[site])
		}
		time.Sleep(delay)
	}
}

func testarSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
