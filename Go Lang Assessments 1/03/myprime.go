package main

import "fmt"

func main() {
	var number, temp int
	temp = 0

	fmt.Print("Enter the Number: ")
	fmt.Scanln(&number)

	for i := 2; i < number/2; i++ {
		if number%i == 0 {
			temp++
			break
		}
	}

	if temp == 0 && number != 1 {
		fmt.Println(number, "is Prime")
	} else {
		fmt.Println(number, "is Not Prime")
	}
}
