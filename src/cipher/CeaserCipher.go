package cipher

import "unicode"

// CaesarCipher представляет структуру для шифрования методом Цезаря.
type CaesarCipher struct{}

// Encrypt шифрует текст с использованием сдвига на заданное количество символов (key).
func (c CaesarCipher) Encrypt(text string, key int) string {
	var result string // Переменная для хранения зашифрованного текста

	// Проходим по каждому символу в строке
	for _, char := range text {
		// Проверяем, является ли символ буквой
		if unicode.IsLetter(char) {
			// Определяем базовое смещение в зависимости от регистра символа
			// Для заглавных букв база будет 'A', для строчных - 'a'
			shiftBase := 'A'
			if unicode.IsLower(char) {
				shiftBase = 'a'
			}

			// Вычисляем новый символ с учетом сдвига
			// Для правильного поведения при сдвиге за пределы диапазона (например, Z -> A или z -> a)
			shiftedChar := (int(char)-int(shiftBase)+key) % 26
			if shiftedChar < 0 {
				shiftedChar += 26 // Учитываем отрицательные сдвиги
			}
			shiftedChar += int(shiftBase)
			
			// Добавляем новый символ к результату
			result += string(rune(shiftedChar))
		} else {
			// Если символ не буква, добавляем его в результат без изменений
			result += string(char)
		}
	}
	return result // Возвращаем зашифрованный текст
}

// Decrypt расшифровывает текст, используя метод Цезаря с отрицательным сдвигом (key).
func (c CaesarCipher) Decrypt(text string, key int) string {
	// Расшифровка осуществляется с использованием шифрования с отрицательным сдвигом
	return c.Encrypt(text, -key)
}
