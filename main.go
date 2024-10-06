package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	// ANSI escape codes for colors used in printing the text
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func main() {
	fmt.Println(Blue + "\n\t| ------------------------------- |")
	fmt.Println("\t| Welcome to the Cypher Tool 2024 |")
	fmt.Println("\t| ------------------------------- |" + Reset)

	// variable for Encrypt or Decrypt choice
	var toEncrypt bool
	// variable for the type of encoding user chooses
	var encoding string
	// variable for the message to be encrypted or decrypted
	var message string

	for {
		toEncrypt, encoding, message = getInput()
		if toEncrypt {
			fmt.Println(Green + "\nEncrypted message using " + Red + encoding + Green + ": " + Reset)
		} else {
			fmt.Println(Green + "\nDecrypted message using " + Red + encoding + Green + ": " + Reset)
		}

		// calls the coresponding encrypting or decrypting function based on users choice from the menu
		switch encoding {
		case "ROT13":
			if toEncrypt {
				fmt.Println(encrypt_rot13(message))
			} else {
				fmt.Println(decrypt_rot13(message))
			}
		case "reverse":
			if toEncrypt {
				fmt.Println(encrypt_reverse(message))
			} else {
				fmt.Println(decrypt_reverse(message))
			}
		case "Random mapping":
			// Custom cipher handling
			if toEncrypt {
				fmt.Println(encrypt_custom(message))
			} else {
				fmt.Println(decrypt_custom(message))
			}
		}

		// Asks user does he want to continue with using the tool
		fmt.Print("\n If you want to continue using the tool choose " + Red + "'Y'" + Reset + " and press \"Enter\": ")
		var noExit string
		reader := bufio.NewReader(os.Stdin)
		noExit, _ = reader.ReadString('\n')
		noExit = strings.TrimSpace(noExit) // Remove newlines and spaces
		if !(noExit == "y" || noExit == "Y") {
			// Exit message
			fmt.Println(Blue + "\n\t****  Thank you for using Cyper Tool 2024  ****\n" + Reset)
			break
		}
	}
}

func getInput() (toEncrypt bool, encoding string, message string) {
	var operationInput string
	var cypherInput string
	toEncrypt = false
	encoding = ""
	// Ввод операции шифрования или дешифрования
	// Input of encryption or decryption operation
	for {
		fmt.Println("\nSelect operation 1 or 2 and press \"Enter\":")
		fmt.Println(Blue + "\t1. Encrypt." + Reset)
		fmt.Println(Blue + "\t2. Decrypt." + Reset)
		fmt.Print(Green + "Enter your input: " + Reset)
		operationInput = readInput()

		if operationInput == "1" {
			toEncrypt = true
			break
		} else if operationInput == "2" {
			toEncrypt = false
			break
		} else {
			fmt.Println(Red + "\n!!! Invalid input. Please select 1 or 2 !!!" + Reset)
		}
	}

	// Ввод выбора шифра
	for {
		fmt.Println("\nSelect cypher and press \"Enter\":")
		fmt.Println(Blue + "\t1. ROT13.")
		fmt.Println("\t2. Reverse.")
		fmt.Println("\t3. Random mapping." + Reset)
		fmt.Print(Green + "Enter your input: " + Reset)
		cypherInput = readInput()

		if cypherInput == "1" {
			encoding = "ROT13"
			break
		} else if cypherInput == "2" {
			encoding = "reverse" // Используем toROT13 для выбора между шифрами
			break
		} else if cypherInput == "3" {
			encoding = "Random mapping"
			break
		} else {
			fmt.Println(Red + "\n!!! Invalid input. Please select 1, 2 or 3 !!!" + Reset)
		}
	}
	// enter message from customer
	for {
		fmt.Println(Green + "\nEnter the message and press \"Enter\": " + Reset)
		message = readInput()
		if message != "" {
			break
		}
		fmt.Println(Red + "\tWrong or empty input! Try again!" + Reset)
	}
	return
}

// readInput reads a line from standard input and returns it as a trimmed string
func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input) // Remove newlines and spaces
}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {
	result := ""
	// temp variable for processing one character at time from the string
	var ch rune
	// with this cycle you go throu all string character by character
	for i := 0; i < len(s); i++ {
		ch = rune(s[i])

		if 'a' <= ch && ch <= 'z' {
			ch += 13
			if ch > 'z' {
				ch -= 26
			}
		}
		// checks is it big letter and after shiftIn is it bigger then 'Z'
		// if its bigger then
		if 'A' <= ch && ch <= 'Z' {
			ch += 13
			if ch > 'Z' {
				ch -= 26
			}
		}
		result += string(ch)
	}
	return result
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {
	// ROT13 decryption is the same as encryption (symmetric)
	return encrypt_rot13(s)
}

// Reverse Alphabet Encryption
func encrypt_reverse(s string) string {
	var result strings.Builder
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			result.WriteRune('Z' - (char - 'A'))
		} else if char >= 'a' && char <= 'z' {
			result.WriteRune('z' - (char - 'a'))
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

// Reverse Alphabet Decryption
func decrypt_reverse(s string) string {
	// Reverse decryption is the same as encryption
	return encrypt_reverse(s)
}

//------------------------------------
// Custom Cipher with map[rune]rune
//------------------------------------

func createEncryptMap() map[rune]rune {
	return map[rune]rune{
		'a': 'M', 'A': 'c', 'b': 'm', 'B': 'Q', 'c': 'S', 'C': 'F',
		'd': 'y', 'D': 'W', 'e': 'q', 'E': 'A', 'f': 'v', 'F': 'X',
		'g': 'N', 'G': 'p', 'h': 'd', 'H': 'u', 'i': 'r', 'I': 'V',
		'j': 'w', 'J': 'Y', 'k': 'g', 'K': 'I', 'l': 'z', 'L': 'T',
		'm': 'j', 'M': 'a', 'n': 'D', 'N': 'K', 'o': 'x', 'O': 't',
		'p': 'Z', 'P': 'B', 'q': 's', 'Q': 'U', 'r': 'C', 'R': 'b',
		's': 'e', 'S': 'E', 't': 'f', 'T': 'G', 'u': 'H', 'U': 'k',
		'v': 'R', 'V': 'P', 'w': 'o', 'W': 'h', 'x': 'l', 'X': 'O',
		'y': 'i', 'Y': 'L', 'z': 'J', 'Z': 'n', '0': '2', '1': '4',
		'2': '6', '3': '1', '4': '8', '5': '7', '6': '0', '7': '9',
		'8': '5', '9': '3',
		'!': ')', '"': '>', '#': '(', '$': '[', '%': '!', '&': '{',
		'\'': ':', '(': ';', ')': '@', '*': '#', '+': ',', ',': '}',
		'-': '~', '.': '"', '/': '\\', ':': '_', ';': '&', '<': '-',
		'=': '.', '>': '`', '?': '*', '@': '%', '[': '|', '\\': '/',
		']': '^', '^': '?', '_': '=', '`': '+', '{': ']', '|': '<',
		'}': '$', '~': '\'',
	}
}

func createDecryptMap() map[rune]rune {
	return map[rune]rune{
		'M': 'a', 'c': 'A', 'm': 'b', 'Q': 'B', 'S': 'c', 'F': 'C',
		'y': 'd', 'W': 'D', 'q': 'e', 'A': 'E', 'v': 'f', 'X': 'F',
		'N': 'g', 'p': 'G', 'd': 'h', 'u': 'H', 'r': 'i', 'V': 'I',
		'w': 'j', 'Y': 'J', 'g': 'k', 'I': 'K', 'z': 'l', 'T': 'L',
		'j': 'm', 'a': 'M', 'D': 'n', 'K': 'N', 'x': 'o', 't': 'O',
		'Z': 'p', 'B': 'P', 's': 'q', 'U': 'Q', 'C': 'r', 'b': 'R',
		'e': 's', 'E': 'S', 'f': 't', 'G': 'T', 'H': 'u', 'k': 'U',
		'R': 'v', 'P': 'V', 'o': 'w', 'h': 'W', 'l': 'x', 'O': 'X',
		'i': 'y', 'L': 'Y', 'J': 'z', 'n': 'Z',
		')': '!', '>': '"', '(': '#', '[': '$', '!': '%', '{': '&',
		':': '\'', ';': '(', '@': ')', '#': '*', ',': '+', '}': ',',
		'~': '-', '"': '.', '\\': '/', '_': ':', '&': ';', '-': '<',
		'.': '=', '`': '>', '*': '?', '%': '@', '|': '[', '/': '\\',
		'^': ']', '?': '^', '=': '_', '+': '`', ']': '{', '<': '|',
		'$': '}', '\'': '~',
	}
}

func encrypt_custom(plainText string) string {
	encryptMap := createEncryptMap()
	var encryptedText strings.Builder

	for _, char := range plainText {
		// First checks is the curren rune is existing in the map list
		// if it does returns the coresponding character from the list
		if _, ok := encryptMap[char]; ok {
			encryptedText.WriteRune(encryptMap[char])
		} else {
			// if the characters is not in the map, then does not encrypt it
			encryptedText.WriteRune(char)
		}
	}
	return encryptedText.String()
}

func decrypt_custom(encryptedText string) string {
	decryptMap := createDecryptMap()
	var decryptedText strings.Builder
	for _, char := range encryptedText {
		// First checks is the curren rune is existing in the map list
		// if it does returns the coresponding character from the list
		if _, ok := decryptMap[char]; ok {
			decryptedText.WriteRune(decryptMap[char])
		} else {
			decryptedText.WriteRune(char)
		}

	}
	return decryptedText.String()
}
