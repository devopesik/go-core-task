package main

import (
	"math/rand"
	"reflect"
	"testing"
)

// Тест генерации случайных чисел
func TestGenerate(t *testing.T) {
	// Зафиксируем seed для воспроизводимости
	rand.Seed(42)

	const size = 5
	const maxValue = 100

	ch := make(chan int)
	go generate(ch, size, maxValue)

	result := make([]int, 0, size)
	for v := range ch {
		if v < 0 || v >= maxValue {
			t.Errorf("Сгенерированное число %d выходит за пределы [0, %d)", v, maxValue)
		}
		result = append(result, v)
	}

	if len(result) != size {
		t.Errorf("Ожидалось %d чисел, получено %d", size, len(result))
	}
}

// Тест: generate закрывает канал
func TestGenerate_ClosesChannel(t *testing.T) {
	ch := make(chan int)
	go generate(ch, 3, 10)

	// Прочитаем все значения
	count := 0
	for range ch {
		count++
		if count > 10 { // защита от бесконечного цикла
			t.Fatal("Канал не закрыт")
		}
	}

	// Если цикл завершился — канал закрыт
	// Дополнительно проверим, что можно безопасно читать из закрытого канала
	if _, ok := <-ch; ok {
		t.Error("Канал не закрыт после генерации")
	}
}

// Тест: size = 0 → возвращает пустой канал
func TestGenerate_ZeroSize(t *testing.T) {
	ch := make(chan int)
	go generate(ch, 0, 100)

	// Канал должен быть сразу закрыт
	_, ok := <-ch
	if ok {
		t.Error("Канал не закрыт при size=0")
	}
}

// Тест: maxValue = 1 → все числа = 0
func TestGenerate_MaxValueOne(t *testing.T) {
	ch := make(chan int)
	go generate(ch, 5, 1)

	result := make([]int, 0, 5)
	for v := range ch {
		result = append(result, v)
	}

	expected := []int{0, 0, 0, 0, 0}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидались нули, получено: %v", result)
	}
}
