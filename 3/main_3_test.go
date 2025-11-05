package main

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

// Тест базовых операций: Add, Get, Exists, Remove
func TestStringIntMap_BasicOperations(t *testing.T) {
	m := NewStringIntMap()

	// Пустая карта
	if m.Exists("key") {
		t.Error("Exists('key') на пустой карте должно быть false")
	}
	if _, ok := m.Get("key"); ok {
		t.Error("Get('key') на пустой карте должно возвращать ok=false")
	}

	// Добавление
	m.Add("a", 1)
	m.Add("b", 2)

	if val, ok := m.Get("a"); !ok || val != 1 {
		t.Errorf("Get('a') = %d, %v; want 1, true", val, ok)
	}
	if !m.Exists("b") {
		t.Error("Exists('b') должно быть true")
	}

	// Обновление
	m.Add("a", 10)
	if val, ok := m.Get("a"); !ok || val != 10 {
		t.Errorf("После обновления Get('a') = %d, %v; want 10, true", val, ok)
	}

	// Удаление
	m.Remove("a")
	if m.Exists("a") {
		t.Error("После Remove('a') Exists('a') должно быть false")
	}
	if _, ok := m.Get("a"); ok {
		t.Error("После Remove('a') Get('a') должно возвращать ok=false")
	}

	// Удаление несуществующего ключа — не должно паниковать
	m.Remove("nonexistent")
}

// Тест Copy — возвращает независимую копию
func TestStringIntMap_Copy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("x", 42)
	m.Add("y", 100)

	copied := m.Copy()

	// Содержимое совпадает
	expected := map[string]int{"x": 42, "y": 100}
	if !reflect.DeepEqual(copied, expected) {
		t.Errorf("Copy() = %v, want %v", copied, expected)
	}

	// Изменение оригинала не влияет на копию
	m.Remove("x")
	if _, exists := copied["x"]; !exists {
		t.Error("Копия не должна меняться при изменении оригинала")
	}

	// Изменение копии не влияет на оригинал
	copied["z"] = 999
	if m.Exists("z") {
		t.Error("Изменение копии не должно влиять на оригинал")
	}
}

// Тест конкурентной безопасности (race condition)
func TestStringIntMap_Concurrent(t *testing.T) {
	const numWorkers = 10
	const numOps = 1000

	m := NewStringIntMap()
	var wg sync.WaitGroup

	// Горутины пишут и читают одновременно
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", workerID%5) // 5 уникальных ключей

			for j := 0; j < numOps; j++ {
				m.Add(key, j)
				m.Exists(key)
				m.Get(key)
				if j%100 == 0 {
					m.Remove("temp") // безопасное удаление несуществующего
				}
			}
		}(i)
	}

	wg.Wait()

	// Проверяем, что данные в валидном состоянии
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		if !m.Exists(key) {
			t.Errorf("Ключ %s должен существовать", key)
		}
	}
}

// Тест String() метода (если реализован)
func TestStringIntMap_String(t *testing.T) {
	m := NewStringIntMap()
	m.Add("hello", 123)
	m.Add("world", 456)

	// Проверяем, что String() возвращает ожидаемую строку
	expected := "map[hello:123 world:456]"
	if got := m.String(); got != expected {
		t.Errorf("String() = %q, want %q", got, expected)
	}
}

// Тест для пустой карты
func TestStringIntMap_Empty(t *testing.T) {
	m := NewStringIntMap()

	if m.Exists("anything") {
		t.Error("Exists на пустой карте должно быть false")
	}

	if _, ok := m.Get("anything"); ok {
		t.Error("Get на пустой карте должно возвращать ok=false")
	}

	copied := m.Copy()
	if len(copied) != 0 {
		t.Errorf("Copy пустой карты должна быть пустой, got %v", copied)
	}

	if m.String() != "map[]" {
		t.Errorf("String() пустой карты = %q, want \"map[]\"", m.String())
	}
}
