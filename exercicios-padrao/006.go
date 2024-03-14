package main

import (
	"fmt"
	"math"
)

func main006() {
	var num float64
	fmt.Print("Digite um número: ")
	fmt.Scanln(&num)
	dobro := num * 2
	triplo := num * 3
	raizq := math.Sqrt(num)
	fmt.Printf("O número é %f, o dobro é %f, o triplo é %f, e a raiz qudrada é %f", num, dobro, triplo, raizq)
}