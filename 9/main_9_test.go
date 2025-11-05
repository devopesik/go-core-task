package main

import (
	"math"
	"reflect"
	"testing"
)

// TestCubeConverter проверяет базовую функциональность
func TestCubeConverter(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	// Запускаем конвертер
	CubeConverter(in, out)

	// Отправляем данные в горутине
	go func() {
		defer close(in)
		for i := uint8(0); i < 5; i++ {
			in <- i
		}
	}()

	// Собираем результаты
	var results []float64
	for val := range out {
		results = append(results, val)
	}

	// Ожидаемые кубы: 0^3, 1^3, ..., 4^3
	expected := []float64{0, 1, 8, 27, 64}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("CubeConverter() = %v, want %v", results, expected)
	}
}

// TestCubeConverter_EmptyInput проверяет пустой вход
func TestCubeConverter_EmptyInput(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	CubeConverter(in, out)

	// Закрываем вход сразу
	close(in)

	// Результат должен быть пустым, и канал out должен закрыться
	var results []float64
	for val := range out {
		results = append(results, val)
	}

	if len(results) != 0 {
		t.Errorf("Expected empty result for empty input, got %v", results)
	}
}

// TestCubeConverter_SingleValue проверяет один элемент
func TestCubeConverter_SingleValue(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	CubeConverter(in, out)

	go func() {
		defer close(in)
		in <- 5
	}()

	results := []float64{}
	for val := range out {
		results = append(results, val)
	}

	expected := []float64{125} // 5^3 = 125
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Single value: got %v, want %v", results, expected)
	}
}

// TestCubeConverter_MaxUint8 проверяет граничное значение
func TestCubeConverter_MaxUint8(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	CubeConverter(in, out)

	go func() {
		defer close(in)
		in <- 255 // max uint8
	}()

	results := []float64{}
	for val := range out {
		results = append(results, val)
	}

	expected := math.Pow(255, 3)
	if len(results) != 1 || results[0] != expected {
		t.Errorf("Max uint8: got %v, want [%v]", results, expected)
	}
}

// TestCubeConverter_Order проверяет сохранение порядка
func TestCubeConverter_Order(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	CubeConverter(in, out)

	input := []uint8{10, 2, 7, 1, 9}
	expected := make([]float64, len(input))
	for i, v := range input {
		expected[i] = math.Pow(float64(v), 3)
	}

	go func() {
		defer close(in)
		for _, v := range input {
			in <- v
		}
	}()

	var results []float64
	for val := range out {
		results = append(results, val)
	}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Order preserved: got %v, want %v", results, expected)
	}
}
