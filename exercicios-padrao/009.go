package main

import "fmt"

func main009() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&num)
	for i := 0; i < 11; i++ {
		var mult int = i * num
		fmt.Printf("%d x %d = %d\n", num, i, mult)
	}
}