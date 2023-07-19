package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для выполнения арифметических операций с арабскими числами
func performArabicOperation(a, b int, operator string) int {
	if operator == "+" {
		return a + b
	} else if operator == "-" {
		return a - b
	} else if operator == "*" {
		return a * b
	} else if operator == "/" {
		return a / b
	} else {
		panic("Недопустимый оператор")
	}
}

// Функция для выполнения арифметических операций с римскими числами
func performRomanOperation(a, b, operator string) (string, error) {
	arabicA, errA := romanToArabic(a)
	arabicB, errB := romanToArabic(b)

	if errA != nil || errB != nil {
		return "", fmt.Errorf("Ошибка: римское число должно быть от I до X")
	}

	// Проверка на диапазон разрешенных римских чисел от I до Х включительно
	if arabicA < 1 || arabicA > 10 || arabicB < 1 || arabicB > 10 {
		return "", fmt.Errorf("Вывод ошибки, так как римские числа не входят в диапазон от 1 до 10 включительно.")
	}

	result := performArabicOperation(arabicA, arabicB, operator)

	if result <= 0 {
		return "", fmt.Errorf("Вывод ошибки, так как в римской системе нет чисел меньше или равных нулю.")
	}

	romanResult := arabicToRoman(result)
	return romanResult, nil
}

// Функция для перевода римского числа в арабское
func romanToArabic(roman string) (int, error) {
	romanNumeralMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
	}

	arabic := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		currentValue, found := romanNumeralMap[string(roman[i])]
		if !found {
			return 0, fmt.Errorf("Ошибка: недопустимый символ в римском числе")
		}

		if currentValue < prevValue {
			arabic -= currentValue
		} else {
			arabic += currentValue
		}
		prevValue = currentValue
	}

	return arabic, nil
}

// Функция для перевода арабского числа в римское
func arabicToRoman(arabic int) string {
	romanNumeralValues := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var sb strings.Builder
	for _, r := range romanNumeralValues {
		for arabic >= r.value {
			sb.WriteString(r.digit)
			arabic -= r.value
		}
	}
	return sb.String()
}

func main() {
	fmt.Println("Введите выражение (например, 1 + 2): ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при считывании данных:", err)
		return
	}

	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	a, operator, b := parts[0], parts[1], parts[2]

	// Проверяем, что используются числа только одной системы счисления
	_, errA := strconv.Atoi(a)
	_, errB := strconv.Atoi(b)

	if (errA == nil && errB == nil) || (errA != nil && errB != nil) {
		if errA == nil && errB == nil {
			// Арабские числа
			arabicA, _ := strconv.Atoi(a)
			arabicB, _ := strconv.Atoi(b)

			// Проверка на числа входящие в диапазон от 1 до 10 включительно
			if arabicA < 1 || arabicA > 10 || arabicB < 1 || arabicB > 10 {
				fmt.Println("Вывод ошибки, так как арабские числа не входят в диапазон от 1 до 10 включительно.")
				return
			}

			result := performArabicOperation(arabicA, arabicB, operator)
			fmt.Println(result)
		} else {
			// Римские числа
			result, err := performRomanOperation(a, b, operator)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		}
	} else {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return
	}
}
