package main

import "fmt"

func main012() {
	var preco float32
	fmt.Print("Insira o Preço: ")
	fmt.Scan(&preco)
	preco5pc := preco * 0.95
	fmt.Printf("O preço com desconto é %f.\n", preco5pc)
}