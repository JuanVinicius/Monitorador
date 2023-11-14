package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const turnos = 3
const delay = 5

func main() {

	intro()

	for {
		menu()
		opcao := comando()

		switch opcao {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
		}
	}
}

func intro() {
	nome := "Juan"
	versao := "1.1.1"
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func comando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func menu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	println("")
	fmt.Println("Iniciando monitorando...")
	sites := []string{"https://www.alura.com.br/", "https://www.google.com.br/"}

	for i := 0; i < turnos; i++ {
		for i, site := range sites {
			fmt.Println("testando site", i, ":", site)
			testaSite(site)
		}
		println("")
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "segue operante ")
	} else {
		fmt.Println("O site:", site, "Está com problemas. Status Code:", resp.StatusCode)
	}
}
