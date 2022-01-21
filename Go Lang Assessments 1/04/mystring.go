package main

import "fmt"

func main() {
	mystr := "JHON\tWEAK"
	fmt.Println("Using tab: ", mystr)

	mystr2 := "Jhon\nweak"
	fmt.Println("using new line: ", mystr2)

	mystr3 := "Mr.ABC"
	fmt.Printf("%q", mystr3)
}
