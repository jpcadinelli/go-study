package main

import "fmt"

func main002() {
	var nome string
	fmt.Print("Digite o seu nome: ")
	fmt.Scanln(&nome)
	fmt.Printf("Olá %s, seja muito bem-vindo!", nome)
}
