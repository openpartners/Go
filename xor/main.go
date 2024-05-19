// xor.exe --cipher --decipher --secret "secret"
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

// var cipherMode = flag.Bool("cipher", true, "Enable cipher mode. This is the default mode.")
// var decipherMode = flag.Bool("decipher", false, "Enable decipher mode.")
// var secretKey = flag.String("secret", "", "You secret key. Must contain at least 1 character")

var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is 'cipher' ")
var secretKey = flag.String("secret", "", "You secret key. Must contain at least 1 character")

func main() {
	flag.Parse()
	// fmt.Println(*cipherMode)

	// if *cipherMode && *decipherMode {
	// 	fmt.Println("Please specify only one mode at a time.") //go run . --cipher --decipher
	// 	os.Exit(1)
	// }

	switch *mode {
	case "cipher":
		plaintext := getUserInput("Enter you text to cipher: ")
		fmt.Println(plaintext)

		cipherer.Cipher(plaintext, *secretKey)
	case "decipher":
		cipheredText := getUserInput("Enter you ciphered data to decipher: ")

		cipherer.Decipher(cipheredText, *secretKey)
	default:
		fmt.Println("Invalid mode. Use 'cipher' or 'decipher'.")
		os.Exit(1)
	}

	if len(*secretKey) == 0 {
		fmt.Println("No secret is provided! Exiting now...")
		os.Exit(1)
	}

}

func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)
	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading the entered text! Please try again")
			continue
		}
		return strings.TrimRight(result, "\n")
	}
}
