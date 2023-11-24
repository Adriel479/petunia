package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
			imprimirLogs()
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
	sites := lerSistesDoArquivo()
	for ii := 0; ii < monitoramentos; ii++ {
		for site := range sites {
			fmt.Println("Testando site", site, ":", sites[site])
			testarSite(sites[site])
		}
		time.Sleep(delay)
	}
}

func testarSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal("Ocorreu um erro ao fazer a requisição: ", err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registrarLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registrarLog(site, false)
	}
}

func lerSistesDoArquivo() []string {
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		log.Fatal("Ocorreu um erro: ", err)
	}
	defer arquivo.Close()
	var (
		leitor = bufio.NewReader(arquivo)
		sites  = make([]string, 0)
	)
	for {
		linha, err := leitor.ReadString('\n')
		if err == io.EOF {
			break
		}
		sites = append(sites, strings.TrimSpace(linha))
	}
	return sites
}

func registrarLog(site string, estado bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Ocorreu um erro:", err)
	}
	defer arquivo.Close()
	arquivo.WriteString(fmt.Sprintf("%s - Site: %s - online: %v\n", time.Now().Format("02/01/2006 15:04:05"), site, estado))
}

func imprimirLogs() {
	fmt.Println("Exibindo logs...")
	arr, err := os.ReadFile("log.txt")
	if err != nil {
		log.Fatal("Ocorreu um erro: ", err)
	}
	fmt.Println(string(arr))
}
