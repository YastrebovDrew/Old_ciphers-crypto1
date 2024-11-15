package cipher

import "strings"

// EnigmaMachine представляет структуру для эмуляции шифрования методом Энигма.
type EnigmaMachine struct {
	rotor1 []rune
	rotor2 []rune
	rotor3 []rune
	reflector []rune
}

// NewEnigmaMachine инициализирует новую машину Энигма с фиксированными роторами и отражателем.
func NewEnigmaMachine() EnigmaMachine {
	// Примерное начальное состояние роторов и отражателя
	return EnigmaMachine{
		rotor1:   []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		rotor2:   []rune("ZYXWVUTSRQPONMLKJIHGFEDCBA"),
		rotor3:   []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		reflector: []rune("YRUHQSLDPXNGOKMIEBFZCWVJAT"),
	}
}

// Encrypt шифрует текст с использованием машины Энигма.
func (e EnigmaMachine) Encrypt(text string) string {
	var encryptedText strings.Builder
	for _, char := range text {
		// Пример простой логики шифрования для символов
		if char >= 'A' && char <= 'Z' {
			// Пропустим через роторы и отражатель (очень упрощенная логика)
			shifted := (int(char-'A') + 3) % 26 // Простая замена с циклическим сдвигом
			encryptedText.WriteRune(rune('A' + shifted))
		} else if char >= 'a' && char <= 'z' {
			shifted := (int(char-'a') + 3) % 26
			encryptedText.WriteRune(rune('a' + shifted))
		} else {
			encryptedText.WriteRune(char) // Пробелы и другие символы остаются без изменений
		}
	}
	return encryptedText.String()
}

// Decrypt расшифровывает текст с использованием машины Энигма.
func (e EnigmaMachine) Decrypt(text string) string {
	var decryptedText strings.Builder
	for _, char := range text {
		// Для простоты предположим, что дешифрование будет просто сдвигом в обратном направлении
		if char >= 'A' && char <= 'Z' {
			shifted := (int(char-'A') - 3 + 26) % 26
			decryptedText.WriteRune(rune('A' + shifted))
		} else if char >= 'a' && char <= 'z' {
			shifted := (int(char-'a') - 3 + 26) % 26
			decryptedText.WriteRune(rune('a' + shifted))
		} else {
			decryptedText.WriteRune(char)
		}
	}
	return decryptedText.String()
}
