package main

import (
	"fmt"
	"strings"
)

// Define rotor mappings (substitution tables)
var rotorI = map[rune]rune{
	'a': 'e', 'b': 'k', 'c': 'm', 'd': 'f', 'e': 'l',
	'f': 'g', 'g': 'd', 'h': 'q', 'i': 'v', 'j': 'z',
	'k': 'n', 'l': 't', 'm': 'o', 'n': 'w', 'o': 'y',
	'p': 'h', 'q': 'x', 'r': 'u', 's': 's', 't': 'p',
	'u': 'a', 'v': 'i', 'w': 'b', 'x': 'r', 'y': 'c',
	'z': 'j',
}

var rotorII = map[rune]rune{
	'a': 'a', 'b': 'j', 'c': 'd', 'd': 'k', 'e': 's',
	'f': 'i', 'g': 'r', 'h': 'u', 'i': 'x', 'j': 'b',
	'k': 'l', 'l': 'h', 'm': 'w', 'n': 't', 'o': 'm',
	'p': 'c', 'q': 'q', 'r': 'g', 's': 'z', 't': 'n',
	'u': 'p', 'v': 'y', 'w': 'f', 'x': 'v', 'y': 'o',
	'z': 'e',
}

var rotorIII = map[rune]rune{
	'a': 'b', 'b': 'd', 'c': 'f', 'd': 'h', 'e': 'j',
	'f': 'l', 'g': 'c', 'h': 'p', 'i': 'r', 'j': 't',
	'k': 'x', 'l': 'v', 'm': 'z', 'n': 'n', 'o': 'y',
	'p': 'e', 'q': 'i', 'r': 'w', 's': 'g', 't': 'a',
	'u': 'k', 'v': 'm', 'w': 'u', 'x': 's', 'y': 'q',
	'z': 'o',
}

func main() {
	plaintext := "hello"
	key := "abc" // Three rotor positions: rotorI, rotorII, and rotorIII

	encryptedText := enigmaEncrypt(plaintext, key)
	decryptedText := enigmaEncrypt(encryptedText, key) // Decryption is the same as encryption in an Enigma machine

	fmt.Println("Plaintext: ", plaintext)
	fmt.Println("Encrypted Text: ", encryptedText)
	fmt.Println("Decrypted Text: ", decryptedText)
}

func enigmaEncrypt(plaintext, key string) string {
	plaintext = strings.ToLower(plaintext)
	var encrypted strings.Builder

	// Initialize rotor positions
	rotorI := []rune("abcdefghijklmnopqrstuvwxyz")
	rotorII := []rune("abcdefghijklmnopqrstuvwxyz")
	rotorIII := []rune("abcdefghijklmnopqrstuvwxyz")

	// Set rotor positions based on the key
	rotorPosition := []int{int(key[0] - 'a'), int(key[1] - 'a'), int(key[2] - 'a')}

	for i, char := range plaintext {
		if char >= 'a' && char <= 'z' {
			// Rotate rotors before encryption
			rotateRotor(&rotorI, rotorPosition[0])
			if i%26 == 0 { // Rotate the second rotor every full rotation of the first rotor
				rotateRotor(&rotorII, 1)
			}
			if i%(26*26) == 0 { // Rotate the third rotor every full rotation of the second rotor
				rotateRotor(&rotorIII, 1)
			}

			// Encrypt the character through the rotors
			encryptedChar := encryptChar(char, rotorI, rotorII, rotorIII)

			// Rotate the first rotor after each character
			rotateRotor(&rotorI, 1)

			encrypted.WriteRune(encryptedChar)
		} else {
			// Non-alphabetic characters are not modified
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}

func rotateRotor(rotor *[]rune, count int) {
	for i := 0; i < count; i++ {
		// Rotate the rotor by one position to the left
		temp := (*rotor)[0]
		copy((*rotor)[:], (*rotor)[1:])
		(*rotor)[len(*rotor)-1] = temp
	}
}

func encryptChar(char rune, rotorI, rotorII, rotorIII []rune) rune {
	// Pass the character through the rotors from right to left
	char = rotorIII[char-'a']
	char = rotorII[char-'a']
	char = rotorI[char-'a']

	// Perform reflection (simple back-and-forth)
	char = rotorI[char-'a']
	char = rotorII[char-'a']
	char = rotorIII[char-'a']

	// Pass the character through the rotors from left to right
	char = rune(strings.IndexRune(string(rotorI), char) + 'a')
	char = rune(strings.IndexRune(string(rotorII), char) + 'a')
	char = rune(strings.IndexRune(string(rotorIII), char) + 'a')

	return char
}
