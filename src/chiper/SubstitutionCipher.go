package cipher

import (
	"math/rand"
	"strings"
	"time"
)

// SubstitutionCipher представляет структуру для шифрования методом подстановки.
type SubstitutionCipher struct {
	upperKey      string // Ключ для подстановки заглавных букв
	lowerKey      string // Ключ для подстановки строчных букв
	upperAlphabet string // Алфавит для заглавных букв
	lowerAlphabet string // Алфавит для строчных букв
}

// NewSubstitutionCipher инициализирует SubstitutionCipher с заданным ключом или случайно сгенерированным.
func NewSubstitutionCipher(key ...string) SubstitutionCipher {
	sc := SubstitutionCipher{
		upperAlphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		lowerAlphabet: "abcdefghijklmnopqrstuvwxyz",
	}

	// Если ключ не передан, генерируем случайный ключ для подстановки
	if len(key) == 0 {
		upperKey := shuffleString(sc.upperAlphabet) // Генерируем случайный ключ для заглавных букв
		lowerKey := shuffleString(sc.lowerAlphabet) // Генерируем случайный ключ для строчных букв
		sc.upperKey = upperKey
		sc.lowerKey = lowerKey
	} else {
		// Если ключ задан, используем его
		sc.upperKey = key[0]
		sc.lowerKey = key[1]
	}

	return sc
}

// shuffleString перемешивает символы строки случайным образом.
func shuffleString(s string) string {
	rand.Seed(time.Now().UnixNano()) // Инициализируем генератор случайных чисел
	runes := []rune(s)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}

// Encrypt шифрует текст, заменяя каждую букву в соответствии с ключом.
func (sc SubstitutionCipher) Encrypt(text string) string {
	// Создаем словарь подстановки для зашифровки
	substitutionDict := make(map[rune]rune)
	for i, char := range sc.upperAlphabet {
		substitutionDict[char] = rune(sc.upperKey[i])
	}
	for i, char := range sc.lowerAlphabet {
		substitutionDict[char] = rune(sc.lowerKey[i])
	}

	// Шифруем текст, заменяя каждый символ по словарю подстановки
	var result strings.Builder
	for _, char := range text {
		if substChar, ok := substitutionDict[char]; ok {
			result.WriteRune(substChar)
		} else {
			result.WriteRune(char) // Если символ не в алфавите, оставляем его без изменений
		}
	}
	return result.String()
}

// Decrypt расшифровывает текст, используя обратную подстановку.
func (sc SubstitutionCipher) Decrypt(text string) string {
	// Создаем словарь подстановки для расшифровки
	substitutionDict := make(map[rune]rune)
	for i, char := range sc.upperKey {
		substitutionDict[char] = rune(sc.upperAlphabet[i])
	}
	for i, char := range sc.lowerKey {
		substitutionDict[char] = rune(sc.lowerAlphabet[i])
	}

	// Расшифровываем текст, заменяя каждый символ по словарю подстановки
	var result strings.Builder
	for _, char := range text {
		if substChar, ok := substitutionDict[char]; ok {
			result.WriteRune(substChar)
		} else {
			result.WriteRune(char) // Если символ не в алфавите, оставляем его без изменений
		}
	}
	return result.String()
}
