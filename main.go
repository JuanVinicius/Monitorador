package main

import (
	"fmt"
	"os"
)

func main() {

	introducao()
	menu()
	opcao := comando()

	switch opcao {
	case 1:
		fmt.Println("Iniciando monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
	}
}

func introducao() {
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
