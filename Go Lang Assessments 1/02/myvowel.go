package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Character: ")
	name, _ := reader.ReadByte()

	if name == 'a' || name == 'e' || name == 'i' || name == 'o' || name == 'u' ||
		name == 'A' || name == 'E' || name == 'I' || name == 'O' || name == 'U' {
		fmt.Printf("%c is a VOWEL Character\n", name)
	}

}
