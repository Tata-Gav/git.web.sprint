package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// ----------------    TO DO list ---------------
	// input has to be trimmed (remove whitespaces from the beginning and the end of the input)
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
	case "randomMap":
		if toEncrypt {
			fmt.Println(encrypt_mapping(message))
		} else {
			fmt.Println(decrypt_mapping(message))
		}

	}

}

func getInput() (toEncrypt bool, encoding string, message string) {
	var operationInput string
	var cypherInput string
	toEncrypt = false
	encoding = ""

	// Ввод операции шифрования или дешифрования-
	// Input of encryption or decryption operation 
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

	// вод выбора шифра
	// ВEntering the cipher selection
	for {
		fmt.Println("\nSelect cypher (1/2):")
		fmt.Println("1. ROT13.")
		fmt.Println("2. Reverse.")
		fmt.Println("3. Random MAP.")
		fmt.Scan(&cypherInput)

		switch cypherInput {
		case "1":
			fmt.Println(" 1 ")
			//Используем toROT13 для выбора между шифрами
			// Using toROT13 to choose between ciphers
			encoding = "ROT13" 
			break
		case "2":
			encoding = "reverse"
			break
		case "3":
			encoding = "randomMap"
			break
		default :
			fmt.Println(" default ")
		}

	}
	
	// enter message from custum
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nEnter the message:")
	message, _ = reader.ReadString('\n')
	// Trim possible spaces at the beginning and end Обрезаем возможные пробелы в начале и конце
	message = message[:len(message)-1] // Remove the newline character Убираем символ новой строки
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
		// it its bigger them
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

//--------------------------------------------------------

// Decrypt the message with rot13
func decrypt_rot13(s string) string {
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
	return encrypt_reverse(s)
}

func encrypt_mapping(s string) string {
	charMap := map[rune]rune{
		'a': 'F', 'B': 'x', 'c': 'Q', 'D': 'l', 'e': 'A', 
		'F': 'm', 'g': 'T', 'H': 'p', 'i': 'W', 'J': 'v', 
		'k': 'U', 'L': 'b', 'm': 'C', 'N': 'r', 'o': 'Y', 
		'P': 'd', 'q': 'G', 'R': 'o', 's': 'K', 'T': 'h', 
		'u': 'V', 'V': 'e', 'w': 'N', 'X': 'j', 'y': 'S', 
		'Z': 'i',
	}

	fmt.Println(charMap['a'])

	var result strings.Builder
	for _, char := range s {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			result.WriteRune(charMap[char])
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()



	//if emoji, exists := emojiMap[char]; exists{}


}

func decrypt_mapping (s string) string {

	// reverseEmoji := make(map[string]rune)
	// for k, v := emojiMap {
	
	return "decrypt mapping not ready"
	//tata
}