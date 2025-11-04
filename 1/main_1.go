package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	// Десятичное число
	var decimalNum int = 100
	fmt.Printf("Тип переменной %T; Десятичное: %d\n", decimalNum, decimalNum)

	// Восьмеричное число
	// Восьмеричное число начинается с нуля (0)
	var octalNum int = 0144 // 0*8^2 + 1*8^1 + 4*8^0 = 64 + 8 + 4 = 100
	fmt.Printf("Тип переменной %T; Восьмеричное: %d\n", octalNum, octalNum)

	// Шестнадцатеричное число
	// Шестнадцатеричное число начинается с 0x
	var hexNum int = 0x64 // 6*16^1 + 4*16^0 = 96 + 4 = 100
	fmt.Printf("Тип переменной %T; Шестнадцатеричное: %d\n", hexNum, hexNum)

	var float64Num float64 = 100
	fmt.Printf("Тип переменной %T; Число с плавающией запятой: %v\n", float64Num, float64Num)

	var stringVar string = "Hello"
	fmt.Printf("Тип переменной %T; Строка: %s\n", stringVar, stringVar)

	var boolVar bool = true
	fmt.Printf("Тип переменной %T; Логическая переменная: %v\n", boolVar, boolVar)

	var complexNum complex64 = 1 + 2i
	fmt.Printf("Тип переменной %T; Комплексная переменная: %v\n", complexNum, complexNum)

	allVars := []any{decimalNum, octalNum, float64Num, hexNum, boolVar, complexNum, stringVar}
	allVarsMerged := merge(allVars)
	fmt.Printf("Тип переменной %T, Значение: %s\n", allVarsMerged, allVarsMerged)

	sliceRunes := []rune(allVarsMerged)
	fmt.Printf("Тип переменной %T, Значение: %v\n", sliceRunes, sliceRunes)

	hashedVar := hash(sliceRunes)
	fmt.Printf("Тип переменной %T, Значение: %v", hashedVar, hashedVar)
}

func merge(values []any) string {
	out := strings.Builder{}

	for _, v := range values {
		stringVal := fmt.Sprintf("%v", v)
		out.WriteString(stringVal)
	}

	return out.String()
}

func hash(runes []rune) [32]byte {
	salt := []byte(string("go-2024"))
	bytes := []byte(string(runes))

	bytes = append(bytes, salt...)
	hashedVar := sha256.Sum256(bytes)

	return hashedVar
}
