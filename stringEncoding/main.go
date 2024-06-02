package main

import (
	"bufio"
	"fmt"
	"os"
)

type RingBuffer struct {
	buffer []byte
	size   int
	read   int
	write  int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]byte, size),
		size:   size,
		read:   0,
		write:  0,
	}
}

func (r *RingBuffer) Write(data byte) {
	r.buffer[r.write] = data
	r.write = (r.write + 1) % r.size
}

func (r *RingBuffer) Read() byte {
	data := r.buffer[r.read]
	r.read = (r.read + 1) % r.size
	return data
}

func Encrypt(input string, shift int) string {
	encrypted := ""
	for i := 0; i < len(input); i++ {
		char := input[i]
		encryptedChar := byte(int(char) + shift)
		encrypted += string(encryptedChar)
	}
	return encrypted
}

func Decrypt(encrypted string, shift int) string {
	decrypted := ""
	for i := 0; i < len(encrypted); i++ {
		char := encrypted[i]
		decryptedChar := byte(int(char) - shift)
		decrypted += string(decryptedChar)
	}
	return decrypted
}

func main() {
	fmt.Println("Введите строку для передачи:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Удаление символа новой строки из входной строки
	input = input[:len(input)-1]

	// Инициализация кольцевого массива с размером, равным длине входной строки
	buffer := NewRingBuffer(len(input))

	// Запись каждого символа в кольцевой массив
	for i := 0; i < len(input); i++ {
		buffer.Write(input[i])
	}

	// Чтение символов из кольцевого массива и шифрование
	encrypted := ""
	for i := 0; i < len(input); i++ {
		char := buffer.Read()
		encrypted += string(char)
	}

	// Расшифровка полученной строки
	decrypted := Decrypt(encrypted, 3)

	fmt.Println("Зашифрованная строка:", encrypted)
	fmt.Println("Расшифрованная строка:", decrypted)
}
