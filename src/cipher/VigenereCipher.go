package cipher

import (
	
	"unicode"
)

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

// extendKey расширяет ключ до длины текста, чтобы использовать его для шифрования и дешифрования.
func (vc VigenereCipher) extendKey(text, key string) string {
	var extendedKey []rune
	keyIndex := 0
	for _, char := range text {
		if unicode.IsLetter(char) { // Проверяем, является ли символ буквой
			extendedKey = append(extendedKey, rune(key[keyIndex%len(key)]))
			keyIndex++ // Увеличиваем индекс ключа только для букв
		} else {
			extendedKey = append(extendedKey, char) // Добавляем небуквенные символы без изменений
		}
	}
	// Выводим расширенный ключ для отладки
	
	return string(extendedKey)
}

// Encrypt шифрует текст с использованием шифра Виженера.
func (vc VigenereCipher) Encrypt(text, key string) string {
	var result []rune
	extendedKey := vc.extendKey(text, key) // Расширяем ключ до длины текста
	for i, char := range text {
		if unicode.IsUpper(char) {
			shiftBase := int('A')
			shift := int(rune(extendedKey[i]) - 'A') // Сдвиг для заглавных
			encryptedChar := rune((int(char)-shiftBase+shift)%26 + shiftBase)
			result = append(result, encryptedChar)
		} else if unicode.IsLower(char) {
			shiftBase := int('a')
			shift := int(rune(extendedKey[i]) - 'a') // Сдвиг для строчных
			encryptedChar := rune((int(char)-shiftBase+shift)%26 + shiftBase)
			result = append(result, encryptedChar)
		} else {
			result = append(result, char) // Оставляем пробелы и символы без изменений
		}
	}
	return string(result)
}

// Decrypt расшифровывает текст с использованием шифра Виженера.
func (vc VigenereCipher) Decrypt(text, key string) string {
	var result []rune
	extendedKey := vc.extendKey(text, key) // Расширяем ключ до длины текста
	for i, char := range text {
		if unicode.IsUpper(char) {
			shiftBase := int('A')
			shift := int(rune(extendedKey[i]) - 'A') // Сдвиг для заглавных
			// Нормализуем сдвиг, чтобы он оставался в пределах от 0 до 25
			decryptedChar := rune((int(char)-shiftBase-shift+26)%26 + shiftBase)
			result = append(result, decryptedChar)
		} else if unicode.IsLower(char) {
			shiftBase := int('a')
			shift := int(rune(extendedKey[i]) - 'a') // Сдвиг для строчных
			// Нормализуем сдвиг, чтобы он оставался в пределах от 0 до 25
			decryptedChar := rune((int(char)-shiftBase-shift+26)%26 + shiftBase)
			result = append(result, decryptedChar)
		} else {
			result = append(result, char) // Оставляем пробелы и символы без изменений
		}
	}
	return string(result)
}
