package main

import "fmt"

func main008() {
	var distm float32
	fmt.Print("Digite uma diatância em metros: ")
	fmt.Scanln(&distm)
	var distcm float32 = distm * 100
	var distmm float32 = distm * 1000
	fmt.Printf("A distância em cm é %f e em mm é %f", distcm, distmm)
}