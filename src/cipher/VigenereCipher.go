package cipher

import "unicode"

// VigenereCipher представляет структуру для шифрования методом Виженера.
type VigenereCipher struct {
	upperAlphabet string // Алфавит заглавных букв
	lowerAlphabet string // Алфавит строчных букв
}

// NewVigenereCipher инициализирует VigenereCipher с алфавитами.
func NewVigenereCipher() VigenereCipher {
	return VigenereCipher{
		upperAlphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		lowerAlphabet: "abcdefghijklmnopqrstuvwxyz",
	}
}

// extendKey расширяет ключ до длины текста.
func (vc VigenereCipher) extendKey(text, key string) string {
	var extendedKey []rune
	keyIndex := 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			extendedKey = append(extendedKey, rune(key[keyIndex%len(key)]))
			keyIndex++
		} else {
			extendedKey = append(extendedKey, char)
		}
	}
	return string(extendedKey)
}

// Encrypt шифрует текст с использованием шифра Виженера.
func (vc VigenereCipher) Encrypt(text, key string) string {
	var result []rune
	extendedKey := vc.extendKey(text, key)

	for i, char := range text {
		if unicode.IsUpper(char) {
			shiftBase := int('A')
			shift := int(rune(extendedKey[i]) - 'A')
			encryptedChar := rune((int(char)-shiftBase+shift)%26 + shiftBase)
			result = append(result, encryptedChar)
		} else if unicode.IsLower(char) {
			shiftBase := int('a')
			shift := int(rune(extendedKey[i]) - 'a')
			encryptedChar := rune((int(char)-shiftBase+shift)%26 + shiftBase)
			result = append(result, encryptedChar)
		} else {
			result = append(result, char)
		}
	}

	return string(result)
}

// Decrypt расшифровывает текст с использованием шифра Виженера.
func (vc VigenereCipher) Decrypt(text, key string) string {
	var result []rune
	extendedKey := vc.extendKey(text, key)

	for i, char := range text {
		if unicode.IsUpper(char) {
			shiftBase := int('A')
			shift := int(rune(extendedKey[i]) - 'A')
			decryptedChar := rune((int(char)-shiftBase-shift+26)%26 + shiftBase)
			result = append(result, decryptedChar)
		} else if unicode.IsLower(char) {
			shiftBase := int('a')
			shift := int(rune(extendedKey[i]) - 'a')
			decryptedChar := rune((int(char)-shiftBase-shift+26)%26 + shiftBase)
			result = append(result, decryptedChar)
		} else {
			result = append(result, char)
		}
	}

	return string(result)
}
