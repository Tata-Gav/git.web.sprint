package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// ----------------    TO DO list ---------------
	// Encrypt the message with reverse
	// Decrypt the message with reverse
	// Encrypt the message with Custom
	// Decrypt the message with Custom
	// Documentation with Markdown

	var toEncrypt bool
	var encoding string
	var message string

	toEncrypt, encoding, message = getInput()

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
	case "custom":
		// Custom cipher handling
		if toEncrypt {
			fmt.Println(encrypt_custom(message))
		} else {
			fmt.Println(decrypt_custom(message))
		}
	}
}

func getInput() (toEncrypt bool, encoding string, message string) {
	var operationInput string
	var cypherInput string
	toEncrypt = false
	encoding = ""

	// Ввод операции шифрования или дешифрования-Input of encryption or decryption operation
	for {
		fmt.Println("\nSelect operation (1/2):")
		fmt.Println("1. Encrypt.")
		fmt.Println("2. Decrypt.")
		fmt.Scan(&operationInput)

		if operationInput == "1" {
			toEncrypt = true
			break
		} else if operationInput == "2" {
			toEncrypt = false
			break
		} else {
			fmt.Println("Invalid input. Please select 1 or 2")
		}
	}

	// Ввод выбора шифра
	for {
		fmt.Println("\nSelect cypher (1/3):")
		fmt.Println("1. ROT13.")
		fmt.Println("2. Reverse.")
		fmt.Println("3. Random mapping.")
		fmt.Scan(&cypherInput)

		if cypherInput == "1" {
			encoding = "ROT13"
			break
		} else if cypherInput == "2" {
			encoding = "reverse" // Используем toROT13 для выбора между шифрами
			break
		} else if cypherInput == "3" {
			encoding = "custom"
			break
		} else {
			fmt.Println("Invalid input. Please select 1, 2 or 3")
		}
	}
	// enter message from custom
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nEnter the message:")
	message, _ = reader.ReadString('\n')
	message = strings.TrimSpace(message) // Remove newlines and spaces
	return
}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {
	result := ""
	var ch rune
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
		'y': 'i', 'Y': 'L', 'z': 'J', 'Z': 'n',
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
	}
}

func encrypt_custom(plainText string) string {
	encryptMap := createEncryptMap()
	var encryptedText strings.Builder

	for _, char := range plainText {
		// First checks is the curren rune is existing in the map list
		// if it does returns and write the coresponding character
		if _, ok := encryptMap[char]; ok {
			encryptedText.WriteRune(encryptMap[char])
		} else {
			encryptedText.WriteRune(char)
		}
	}

	return encryptedText.String()
}

func decrypt_custom(encryptedText string) string {
	decryptMap := createDecryptMap()
	var decryptedText strings.Builder
	for _, char := range encryptedText {
		if _, ok := decryptMap[char]; ok {
			decryptedText.WriteRune(decryptMap[char])
		} else {
			decryptedText.WriteRune(char)
		}

	}
	return decryptedText.String()
}
