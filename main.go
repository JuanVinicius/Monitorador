package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {
		intro()
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
	fmt.Println("Iniciando monitorando...")
	site := "https://www.alura.com.br/"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "segue operante ")
	} else {
		fmt.Println("O site:", site, "Está com problemas. Status Code:", resp.StatusCode)
	}
}
