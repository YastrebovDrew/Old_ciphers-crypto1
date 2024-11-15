package cipher

import (
	"math/rand"
	"strings"
	"time"
)

// SubstitutionCipher представляет структуру для шифрования методом подстановки.
type SubstitutionCipher struct {
	upperKey      string
	lowerKey      string
	upperAlphabet string
	lowerAlphabet string
}

// NewSubstitutionCipher инициализирует SubstitutionCipher с заданным ключом или случайно сгенерированным.
func NewSubstitutionCipher(key ...string) SubstitutionCipher {
	sc := SubstitutionCipher{
		upperAlphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		lowerAlphabet: "abcdefghijklmnopqrstuvwxyz",
	}

	if len(key) == 0 {
		// Генерация случайных ключей для подстановки
		sc.upperKey = shuffleString(sc.upperAlphabet)
		sc.lowerKey = shuffleString(sc.lowerAlphabet)
	} else {
		// Использование переданных ключей
		sc.upperKey = key[0]
		sc.lowerKey = key[1]
	}

	return sc
}

// shuffleString перемешивает символы строки случайным образом.
func shuffleString(s string) string {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	runes := []rune(s)
	rng.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}

// Encrypt шифрует текст с использованием подстановки.
func (sc SubstitutionCipher) Encrypt(text string) string {
	substitutionDict := make(map[rune]rune)
	// Подстановка для заглавных букв
	for i, char := range sc.upperAlphabet {
		substitutionDict[char] = rune(sc.upperKey[i])
	}
	// Подстановка для строчных букв
	for i, char := range sc.lowerAlphabet {
		substitutionDict[char] = rune(sc.lowerKey[i])
	}

	var result strings.Builder
	for _, char := range text {
		if substChar, ok := substitutionDict[char]; ok {
			result.WriteRune(substChar)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

// Decrypt расшифровывает текст с использованием подстановки.
func (sc SubstitutionCipher) Decrypt(text string) string {
	substitutionDict := make(map[rune]rune)
	// Подстановка для заглавных букв
	for i, char := range sc.upperKey {
		substitutionDict[char] = rune(sc.upperAlphabet[i])
	}
	// Подстановка для строчных букв
	for i, char := range sc.lowerKey {
		substitutionDict[char] = rune(sc.lowerAlphabet[i])
	}

	var result strings.Builder
	for _, char := range text {
		if substChar, ok := substitutionDict[char]; ok {
			result.WriteRune(substChar)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
