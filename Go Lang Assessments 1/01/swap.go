package main

import "fmt"

func main() {
	num1, num2 := 10, 20
	fmt.Println("num1 is: ", num1)
	fmt.Println("num2 is: ", num2)

	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2

	fmt.Println("After swap..")
	fmt.Printf("num1 is %d and num2 is: %d", num1, num2)
}
