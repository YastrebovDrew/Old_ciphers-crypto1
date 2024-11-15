package main

import (
    "fmt"
    "your_module_name/cipher" // Импортируем пакет с шифрами
)

func main() {
    plaintext := "Hello, my name is Kalashov Feodor Olegovich and I am a 4th year student of group N34481"

    // Шифр Цезаря
    caesarCipher := cipher.CaesarCipher{}
    shift := 3

    encryptedText := caesarCipher.Encryption(plaintext, shift)
    fmt.Println("Зашифрованный текст шифра Цезаря:", encryptedText)

    decryptedText := caesarCipher.Decryption(encryptedText, shift)
    fmt.Println("Расшифрованный текст шифра Цезаря:", decryptedText)

    fmt.Println("\n")

    // Моноалфавитный шифр замены
    substitutionCipher := cipher.SubstitutionCipher{}

    substitutionEncrypted := substitutionCipher.Encryption(plaintext)
    fmt.Println("Зашифрованный текст шифра замены:", substitutionEncrypted)

    substitutionDecrypted := substitutionCipher.Decryption(substitutionEncrypted)
    fmt.Println("Расшифрованный текст шифра замены:", substitutionDecrypted)

    fmt.Println("\n")

    // Моноалфавитный шифр перестановки
    transpositionCipher := cipher.TranspositionCipher{}
    transpositionKey := "4312"

    encryptedText = transpositionCipher.Encryption(plaintext, transpositionKey)
    fmt.Println("Зашифрованный текст шифра перестановки:", encryptedText)

    decryptedText = transpositionCipher.Decryption(encryptedText, transpositionKey)
    fmt.Println("Расшифрованный текст шифра перестановки:", decryptedText)

    fmt.Println("\n")

    vigenereCipher := cipher.VigenereCipher{}
    vigenereKey := "LEMON"

    encryptedText = vigenereCipher.Encryption(plaintext, vigenereKey)
    fmt.Println("Зашифрованный текст шифра Виженера:", encryptedText)

    decryptedText = vigenereCipher.Decryption(encryptedText, vigenereKey)
    fmt.Println("Расшифрованный текст шифра Виженера:", decryptedText)

    fmt.Println("\n")

    // Настройки машины "Энигма"
    rotor1 := cipher.Rotor{Wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ", Notch: 16}
    rotor2 := cipher.Rotor{Wiring: "AJDKSIRUXBLHWTMCQGZNPYFVOE", Notch: 4}
    rotor3 := cipher.Rotor{Wiring: "BDFHJLCPRTXVZNYEIWGAKMUSQO", Notch: 21}
    reflector := cipher.Reflector{Wiring: "YRUHQSLDPXNGOKMIEBFZCWVJAT"}

    // Создание машины и шифрование
    enigma := cipher.EnigmaMachine{Rotors: []cipher.Rotor{rotor1, rotor2, rotor3}, Reflector: reflector}
    message := "HELLOMYNAMEISKALASHOVFEODOROLEGOVICH"
    encryptedMessage := enigma.EncodeMessage(message)
    fmt.Println("Зашифрованное сообщение:", encryptedMessage)

    // Сброс позиции роторов перед дешифрованием
    enigma.Rotors[0].Position = 0
    enigma.Rotors[1].Position = 0
    enigma.Rotors[2].Position = 0

    // Дешифрование сообщения
    decryptedMessage := enigma.EncodeMessage(encryptedMessage)
    fmt.Println("Расшифрованное сообщение:", decryptedMessage)
}
