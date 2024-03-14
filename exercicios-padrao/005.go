package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número intero: ")
	fmt.Scan(&num)
	var ant int = num - 1
	var suc int = num + 1
	fmt.Printf("Seu antecessor é %d e seu sucessor é %d", ant, suc)
}