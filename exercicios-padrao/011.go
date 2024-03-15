package main

import "fmt"

func main011() {
	var largura float32
	var altura float32
	fmt.Println("Qual a largura e a altura da sua parede em metros?")
	fmt.Scan(&largura, &altura)
	areaParede := largura * altura
	litrosTinta := areaParede / 2
	fmt.Printf("Você vai precisar de %f Litros de tinta.\n", litrosTinta)
}