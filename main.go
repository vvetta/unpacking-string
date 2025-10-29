package main

import (
	"bufio"
	"fmt"
	"os"

	unpackingstring "github.com/vvetta/unpacking-string/unpackingString"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var inputString string
	fmt.Print("Введите строку, которую хотите раcпаковать: ")
	fmt.Fscan(reader, &inputString)

	result, err := unpackingstring.Unpack(inputString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	fmt.Println("Result: ", result)
}
