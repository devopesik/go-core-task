package main

import (
	"reflect"
	"testing"
)

// Тест: объединение двух непустых каналов
func TestMerge_TwoChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Заполняем каналы в горутинах
	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 10; i <= 12; i++ {
			ch2 <- i
		}
	}()

	resultChan := merge(ch1, ch2)

	var result []int
	for v := range resultChan {
		result = append(result, v)
	}

	expected := []int{1, 2, 3, 10, 11, 12}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("merge() = %v, want %v", result, expected)
	}
}

// Тест: один канал пустой
func TestMerge_OneEmptyChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		// ничего не отправляем
	}()

	go func() {
		defer close(ch2)
		ch2 <- 42
		ch2 <- 43
	}()

	resultChan := merge(ch1, ch2)

	var result []int
	for v := range resultChan {
		result = append(result, v)
	}

	expected := []int{42, 43}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("merge() = %v, want %v", result, expected)
	}
}

// Тест: все каналы пустые
func TestMerge_AllEmpty(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() { close(ch1) }()
	go func() { close(ch2) }()

	resultChan := merge(ch1, ch2)

	var result []int
	for v := range resultChan {
		result = append(result, v)
	}

	if len(result) != 0 {
		t.Errorf("merge() = %v, want empty slice", result)
	}
}

// Тест: один канал
func TestMerge_SingleChannel(t *testing.T) {
	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 100
		ch <- 200
	}()

	resultChan := merge(ch)

	var result []int
	for v := range resultChan {
		result = append(result, v)
	}

	expected := []int{100, 200}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("merge() = %v, want %v", result, expected)
	}
}

// Тест: три канала
func TestMerge_ThreeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
	}()

	go func() {
		defer close(ch2)
		ch2 <- 2
	}()

	go func() {
		defer close(ch3)
		ch3 <- 3
	}()

	resultChan := merge(ch1, ch2, ch3)

	var result []int
	for v := range resultChan {
		result = append(result, v)
	}

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("merge() = %v, want %v", result, expected)
	}
}

// Тест: закрытие выходного канала
func TestMerge_ClosesOutput(t *testing.T) {
	ch := make(chan int)
	go func() { close(ch) }()

	out := merge(ch)

	// Должен завершиться сразу
	_, ok := <-out
	if ok {
		t.Error("Выходной канал не закрыт")
	}
}
