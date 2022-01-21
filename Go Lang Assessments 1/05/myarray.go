package main

import "fmt"

func main() {
	myarray := [11]int{2, 4, 6, 8, 34, 5, 7, 3, 5, 67, 4}
	var sum int = 0
	for i := 0; i < len(myarray); i++ {
		sum = sum + myarray[i]
	}
	fmt.Println("sum is: ", sum)
	avg := sum / len(myarray)
	fmt.Println("average: ", avg)
}
