package main

import "fmt"

func main007() {
	var nota1 float32
	var nota2 float32
	fmt.Print("Digite 2 notas: ")
	fmt.Scan(&nota1, &nota2)
	media := (nota1 + nota2) / 2
	fmt.Printf("A média das notas é %f", media)
}