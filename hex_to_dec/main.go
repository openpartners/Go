package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println("Enter hex number or 'stop' to exit:")

	// var input, i2, i3 string
	// fmt.Scanln(&input, &i2, &i3)
	// fmt.Println(input, i2, i3)
	// fmt.Println(input)
	var input string
	fmt.Scanln(&input)

	input = strings.ToLower(input)

	if input == "stop" {
		//...
	}

	i := new(big.Int)

	convertedValue, ok := i.SetString(processHex(input), 16) // 0xabc123
	if !ok {
		fmt.Println("failed!")
	} else {
		fmt.Println(convertedValue, i)
	}
}

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}
