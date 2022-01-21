package main

import "fmt"

func MyFact(number int) uint64 {
	if number >= 1 {
		return uint64(number) * MyFact(number-1)
	} else {
		return 1
	}
}
func main() {
	var number int
	fmt.Print("Enter Number:")
	fmt.Scanln(&number)
	fmt.Printf("Factorial of %d : %d\n", number, MyFact(number))

}
