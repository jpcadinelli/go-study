package main

import "fmt"

func main013() {
	var salario float32
	fmt.Print("Insira o salário: ")
	fmt.Scan(&salario)
	reajuste := salario * 1.15
	fmt.Printf("O salário com reajuste é R$%f\n", reajuste)
}