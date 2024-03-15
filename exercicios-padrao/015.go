package main

import "fmt"

func main015() {
	var distkm float32
	var dias float32
	fmt.Print("Insira a quantidade de Km e a quantidade de dias: ")
	fmt.Scan(&distkm, &dias)
	valor := distkm * 0.15 + dias * 60
	fmt.Printf("O valor de pagamento do aluguel do carro é R$%.2f\n", valor)
}