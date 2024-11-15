package cipher

import (
    "strings"
    "sort"
)

// TranspositionCipher представляет структуру для шифрования методом перестановки.
type TranspositionCipher struct{}

// Encrypt шифрует текст с использованием перестановки по заданному ключу.
func (tc TranspositionCipher) Encrypt(text string, key string) string {
    n := len(key)
    // Дополняем текст пробелами до кратности длине ключа
    paddingSize := n - len(text)%n
    if paddingSize < n {
        text += strings.Repeat(" ", paddingSize)
    }

    // Разбиваем текст на части длиной n
    chunks := make([]string, 0, len(text)/n)
    for i := 0; i < len(text); i += n {
        chunks = append(chunks, text[i:i+n])
    }

    // Сортируем ключ и создаем индексы для перестановки
    sortedKey := strings.Split(key, "")
    sort.Strings(sortedKey)
    indexOrder := make([]int, n)
    for i, char := range sortedKey {
        indexOrder[i] = strings.Index(key, char)
    }

    // Переставляем символы в каждом фрагменте в соответствии с порядком индексов
    var encryptedText strings.Builder
    for _, chunk := range chunks {
        for _, i := range indexOrder {
            encryptedText.WriteByte(chunk[i])
        }
    }

    return encryptedText.String()
}

// Decrypt расшифровывает текст, используя обратную перестановку по ключу.
func (tc TranspositionCipher) Decrypt(text string, key string) string {
    n := len(key)

    // Разбиваем зашифрованный текст на части длиной n
    chunks := make([]string, 0, len(text)/n)
    for i := 0; i < len(text); i += n {
        chunks = append(chunks, text[i:i+n])
    }

    // Сортируем ключ и создаем индексы для обратной перестановки
    sortedKey := strings.Split(key, "")
    sort.Strings(sortedKey)
    indexOrder := make([]int, n)
    for i, char := range sortedKey {
        indexOrder[i] = strings.Index(key, char)
    }

    // Создаем обратный порядок индексов для дешифровки
    inverseIndexOrder := make([]int, n)
    for i, index := range indexOrder {
        inverseIndexOrder[index] = i
    }

    // Восстанавливаем исходный текст
    var decryptedText strings.Builder
    for _, chunk := range chunks {
        decryptedChunk := make([]byte, n)
        for i, pos := range inverseIndexOrder {
            decryptedChunk[pos] = chunk[i]
        }
        decryptedText.WriteString(string(decryptedChunk))
    }

    return strings.TrimRight(decryptedText.String(), " ") // Убираем добавленные пробелы
}
