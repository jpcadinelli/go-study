package main

import "fmt"

func main014() {
	var tempc float32
	fmt.Print("Insira a temperatura em °C: ")
	fmt.Scan(&tempc)
	tempf := ((9 * tempc) / 5) +32
	fmt.Printf("A temperatura %f°C é %f°F\n", tempc, tempf)
}