package cipher

import (
	"strings"
)

// Rotor представляет ротор Enigma с проводкой, позицией и щелчком
type Rotor struct {
	wiring   string // Проводка ротора
	notch    int    // Позиция щелчка, при которой ротор поворачивает следующий ротор
	position int    // Текущая позиция ротора
}

// NewRotor создает новый ротор с заданной проводкой, щелчком и начальной позицией
func NewRotor(wiring string, notch, position int) *Rotor {
	return &Rotor{wiring: wiring, notch: notch, position: position}
}

// EncodeForward кодирует символ при движении вперед через ротор
func (r *Rotor) EncodeForward(char rune) rune {
	// Преобразуем символ с учетом текущей позиции ротора
	index := (int(char) - 'A' + r.position) % 26
	encodedChar := r.wiring[index]
	return rune((int(encodedChar)-'A'-r.position+26)%26 + 'A')
}

// EncodeBackward кодирует символ при обратном движении через ротор
func (r *Rotor) EncodeBackward(char rune) rune {
	// Ищем позицию символа в проводке для обратного кодирования
	index := (strings.IndexRune(r.wiring, rune((int(char)-'A'+r.position)%26+'A')) - r.position + 26) % 26
	return rune(index + 'A')
}

// Rotate поворачивает ротор, возвращая true, если достигнута позиция щелчка
func (r *Rotor) Rotate() bool {
	r.position = (r.position + 1) % 26
	return r.position == r.notch
}

// Reflector представляет отражатель Enigma с проводкой
type Reflector struct {
	wiring string // Проводка отражателя
}

// NewReflector создает новый отражатель с заданной проводкой
func NewReflector(wiring string) *Reflector {
	return &Reflector{wiring: wiring}
}

// Reflect отражает символ, перенаправляя его обратно через роторы
func (rf *Reflector) Reflect(char rune) rune {
	return rune(rf.wiring[int(char)-'A'])
}

// EnigmaMachine представляет машину Enigma с роторами и отражателем
type EnigmaMachine struct {
	rotors    []*Rotor
	reflector *Reflector
}

// NewEnigmaMachine создает новую машину Enigma с роторами и отражателем
func NewEnigmaMachine(rotors []*Rotor, reflector *Reflector) *EnigmaMachine {
	return &EnigmaMachine{rotors: rotors, reflector: reflector}
}

// EncodeCharacter шифрует символ, пропуская его через роторы, отражатель и снова через роторы
func (em *EnigmaMachine) EncodeCharacter(char rune) rune {
	// Прямое прохождение через роторы
	for _, rotor := range em.rotors {
		char = rotor.EncodeForward(char)
	}

	// Отражение
	char = em.reflector.Reflect(char)

	// Обратное прохождение через роторы
	for i := len(em.rotors) - 1; i >= 0; i-- {
		char = em.rotors[i].EncodeBackward(char)
	}

	return char
}

// StepRotors поворачивает роторы, учитывая щелчки
func (em *EnigmaMachine) StepRotors() {
	rotateNext := em.rotors[0].Rotate() // Вращаем первый ротор
	if rotateNext {
		rotateNext = em.rotors[1].Rotate() // Вращаем второй ротор, если первый достиг щелчка
		if rotateNext {
			em.rotors[2].Rotate() // Вращаем третий ротор, если второй достиг щелчка
		}
	}
}

// EncodeMessage шифрует сообщение, по одному символу за раз
func (em *EnigmaMachine) EncodeMessage(message string) string {
	encryptedMessage := ""
	for _, char := range message {
		em.StepRotors()                       // Вращаем роторы перед шифрованием символа
		encryptedMessage += string(em.EncodeCharacter(char)) // Шифруем символ
	}
	return encryptedMessage
}
