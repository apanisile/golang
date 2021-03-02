package main

import "fmt"

func main() {
	//Go program to convert celsius to Fahrenheit
	fmt.Print("Enter the value in Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	celsius := (input - 32) * 5 / 9

	fmt.Println(celsius)

}
