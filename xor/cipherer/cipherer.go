package cipherer

import (
	"encoding/base64"
	"fmt"
	"os"
)

func Cipher(rawString, secret string) string {
	return base64.StdEncoding.EncodeToString(
		process(
			[]byte(rawString),
			[]byte(secret),
		),
	)
}

func Decipher(cipheredText, secret string) string {
	cipheredBytes, err := base64.StdEncoding.DecodeString(cipheredText)

	if err != nil {
		fmt.Println("An error occured while processing the ciphered data! Exiting now...")
		os.Exit(1)
	}

	decriptedBytes := process(cipheredBytes, []byte(secret))

	if len(decriptedBytes) == 0 {
		fmt.Println("Dectription failed due to invalid input or secret. Exiting now...")
		os.Exit(1)
	}

	return string(decriptedBytes)
}

func process(input, secret []byte) []byte {
	if len(secret) == 0 {
		fmt.Println("Secret cannot be empty. Exiting now...")
		os.Exit(1)
	}

	for i, b := range input {
		input[i] = b ^ secret[i%len(secret)] // 0..len(secret)
	}
	return input
}
