package main

import (
	"fmt"

	"old_ciphers/cipher" // Поменяйте на ваш правильный путь
)

func main() {
	// Пример использования шифров
	text := "Hello, my name is Yastrebov Andrey!"

	// Substitution Cipher
	substitution := cipher.NewSubstitutionCipher() // Генерация случайного ключа
	encryptedSub := substitution.Encrypt(text)
	decryptedSub := substitution.Decrypt(encryptedSub)
	fmt.Println("Substitution Cipher:")
	fmt.Println("Encrypted:", encryptedSub)
	fmt.Println("Decrypted:", decryptedSub)

	// Transposition Cipher
	transposition := cipher.TranspositionCipher{}
	key := "KEY"
	encryptedTrans := transposition.Encrypt(text, key)
	decryptedTrans := transposition.Decrypt(encryptedTrans, key)
	fmt.Println("\nTransposition Cipher:")
	fmt.Println("Encrypted:", encryptedTrans)
	fmt.Println("Decrypted:", decryptedTrans)

	// Vigenere Cipher
	vigenere := cipher.NewVigenereCipher()
	keyVigenere := "KEY"
	encryptedVigenere := vigenere.Encrypt(text, keyVigenere)
	decryptedVigenere := vigenere.Decrypt(encryptedVigenere, keyVigenere)
	fmt.Println("\nVigenere Cipher:")
	fmt.Println("Encrypted:", encryptedVigenere)
	fmt.Println("Decrypted:", decryptedVigenere)

	// Enigma Machine
	enigma := cipher.NewEnigmaMachine() // Предположим, что у вас есть такая функция
	encryptedEnigma := enigma.Encrypt(text)
	decryptedEnigma := enigma.Decrypt(encryptedEnigma)
	fmt.Println("\nEnigma Machine:")
	fmt.Println("Encrypted:", encryptedEnigma)
	fmt.Println("Decrypted:", decryptedEnigma)

	// Пример работы с ошибками
	// Если вы хотите обработать возможные ошибки, то используйте `log.Fatal` для остановки
	// и обработки ошибок
	err := simulateError() // Функция, которая может генерировать ошибку
	if err != nil {
	}
}

// simulateError - пример функции, которая может вызвать ошибку.
func simulateError() error {
	return fmt.Errorf("simulated error")
}
