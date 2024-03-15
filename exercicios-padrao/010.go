package main

import "fmt"

func main010() {
	var saldo float32
	fmt.Print("Insira quanto dinheiro você tem na conta: ")
	fmt.Scanln(&saldo)
	dolares := saldo / 3.27
	fmt.Printf("Você pode comprar %f dolares\n", dolares)
}