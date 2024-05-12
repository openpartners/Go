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

	for {
		fmt.Scanln(&input)

		input = strings.ToLower(input)

		if input == "stop" {
			break
		}

		i := new(big.Int)

		// convertedValue, ok := i.SetString(processHex(input), 16) // 0xabc123
		// _, ok := i.SetString(processHex(input), 16) // 0xabc123

		// if !ok {
		// 	fmt.Println("failed!")
		// } else {
		// 	// fmt.Println(convertedValue, i)
		// 	fmt.Println(i)
		// }
		if _, ok := i.SetString(processHex(input), 16); !ok {
			fmt.Println("Invalid hexadecimal number!")
			continue // next iteration - for
		}
		fmt.Println(i)
	}
}

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}
