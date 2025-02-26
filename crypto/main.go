package main

import (
	"bufio"
	"fmt"
	"os"

	"crypto/rand"

	// "github.com/tink-crypto/tink-go/v2/keyset"
	// "github.com/tink-crypto/tink-go/v2/streamingaead"
	"golang.org/x/crypto/argon2"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var password string
	for {
		print("Enter your super cool password: ")

		scanner.Scan()
		text := scanner.Text()

		if len(text) != 0 {
			password = text
			break
		} else if len(text) > 72 {
			println("Password exceeds maximum length of 72 characters.")
		}
	}
	println()

	// keysetHandle, err := keyset.NewHandle(streamingaead.AES128GCMHKDF1MBKeyTemplate())
	// if err != nil {
	// 	panic(err)
	// }

	// keysetManager := keyset.NewManagerFromHandle(keysetHandle)
	// println("Keyset Manager:")
	// println(keysetManager)

	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	fmt.Printf("Salt: %x\n", salt)

	derivedKey := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 16)
	fmt.Printf("Derived Key: %x\n", derivedKey)

	os.Exit(0)
}
