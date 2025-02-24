package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tink-crypto/tink-go/v2/streamingaead"
	"golang.org/x/crypto/argon2"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	userPassword := ""
	for {
		fmt.Print("Enter your really cool password (72 chars max): ")
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) == 0 {
			panic("user password not specified")
		} else if len(text) > 72 {
			panic("user password longer than 72 characters")
		} else {
			userPassword = text
			break
		}
	}

	encryptionKey := argon2.IDKey([]byte(userPassword), []byte("saltysalt"), 1, 64*1024, 4, 32)

	streamingaead.New()

	// cipher := adiantum.New(encryptionKey)
	// tweak := make([]byte, 12) // can be any length
	// plaintext := []byte("Hello, world! This is a bit longer than 16 bytes.")
	// ciphertext := cipher.Encrypt(plaintext, tweak)
	// recovered := cipher.Decrypt(ciphertext, tweak)
	// println(string(recovered)) // Hello, world!
}
