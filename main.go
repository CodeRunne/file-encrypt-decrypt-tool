package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file and decrypt to decrypt")
		os.Exit(0)
	}
}

func printHelp() {
	fmt.Println("File encryption")
	fmt.Println("Simple file encrypter for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tgo run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTries to decrypts a file given a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. For more info, run go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	password := getPassword()
	fmt.Println("\nEncrypting...")
	Encrypt(file, password)
	fmt.Println("\n file successfully protected")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. For more info, run go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	fmt.Print("Enter password: ")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nDecrypting...")
	Decrypt(file, password)
	fmt.Println("\nFile successfully decrypted")
}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm Password: ")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\nPasswords do not match. Please try again")
		return getPassword()
	}
	return password
}

func validatePassword(password1, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	} else {
		return true
	}
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
