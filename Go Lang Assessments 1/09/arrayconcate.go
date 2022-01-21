package main

import "fmt"

func main() {
	array1 := []int{3, 5, 76, 3, 6, 3, 5, 6, 3}
	array2 := []int{2, 3, 65, 7, 4, 3, 6, 3, 56, 3}
	resultarr := append([]int{}, append(array1, array2...)...)
	fmt.Println(resultarr)
}
