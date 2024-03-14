package main

import "fmt"

func main003() {
	var num1 int
	var num2 int
	fmt.Println("Digite 2 números: ")
	fmt.Scan(&num1, &num2)
	sum := num1 + num2
	fmt.Printf("A soma é %d", sum)
}