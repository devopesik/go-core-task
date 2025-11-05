package main

import (
	"reflect"
	"testing"
)

// Тест для generateRandomSlice
func TestGenerateRandomSlice(t *testing.T) {
	const size, max = 10, 100

	// Запускаем несколько раз, чтобы убедиться в стабильности
	for i := 0; i < 10; i++ {
		result := generateRandomSlice(size, max)

		// Проверяем длину
		if len(result) != size {
			t.Errorf("Expected length %d, got %d", size, len(result))
		}

		// Проверяем уникальность
		seen := make(map[int]bool)
		for _, v := range result {
			if v < 0 || v >= max {
				t.Errorf("Value %d is out of range [0, %d)", v, max)
			}
			if seen[v] {
				t.Errorf("Duplicate value found: %d", v)
			}
			seen[v] = true
		}
	}
}

func TestGenerateRandomSlice_InvalidSize(t *testing.T) {
	result := generateRandomSlice(5, 5) // [0,1,2,3,4] в случайном порядке
	if len(result) != 5 {
		t.Fatalf("Expected 5 elements, got %d", len(result))
	}

	seen := make(map[int]bool)
	for _, v := range result {
		if v < 0 || v >= 5 {
			t.Errorf("Value %d out of range [0,5)", v)
		}
		seen[v] = true
	}
	if len(seen) != 5 {
		t.Errorf("Not all values from 0 to 4 are present")
	}
}

// Тест для sliceExample (фильтрация чётных)
func TestSliceExample(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"all even", []int{2, 4, 6}, []int{2, 4, 6}},
		{"all odd", []int{1, 3, 5}, []int{}},
		{"mixed", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"empty", []int{}, []int{}},
		{"with zero", []int{0, -2, 3}, []int{0, -2}}, // 0 и отрицательные чётные
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.input)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("sliceExample(%v) = %v, want %v", tt.input, got, tt.expect)
			}
		})
	}
}

// Тест для addElements
func TestAddElements(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		num    int
		expect []int
	}{
		{"non-empty", []int{1, 2, 3}, 99, []int{1, 2, 3, 99}},
		{"empty", []int{}, 42, []int{42}},
		{"single", []int{5}, 0, []int{5, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElements(tt.slice, tt.num)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("addElements(%v, %d) = %v, want %v", tt.slice, tt.num, got, tt.expect)
			}

			// Проверяем, что исходный срез не изменился
			if len(tt.slice) > 0 && &got[0] == &tt.slice[0] {
				t.Error("addElements modified the original slice (shallow copy)")
			}
		})
	}
}

// Тест для copySlice
func TestCopySlice(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	copied := copySlice(original)

	if !reflect.DeepEqual(original, copied) {
		t.Errorf("copySlice failed: got %v, want %v", copied, original)
	}

	// Проверяем, что это разные слайсы (разные адреса)
	if len(original) > 0 && &original[0] == &copied[0] {
		t.Error("copySlice returned a shallow copy (same underlying array)")
	}

	// Проверяем изменение копии не влияет на оригинал
	copied[0] = 999
	if original[0] == 999 {
		t.Error("Modifying copy affected original slice")
	}
}

// Тест для removeElement
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name        string
		slice       []int
		index       int
		expect      []int
		shouldPanic bool
	}{
		{"middle", []int{1, 2, 3, 4}, 2, []int{1, 2, 4}, false},
		{"first", []int{1, 2, 3}, 0, []int{2, 3}, false},
		{"last", []int{1, 2, 3}, 2, []int{1, 2}, false},
		{"single", []int{42}, 0, []int{}, false},
		{"empty", []int{}, 0, nil, true},             // паника: индекс за пределами
		{"out of bounds", []int{1, 2}, 5, nil, true}, // паника
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if tt.shouldPanic {
					if r == nil {
						t.Error("Expected panic, but none occurred")
					}
				} else {
					if r != nil {
						t.Errorf("Unexpected panic: %v", r)
					}
				}
			}()

			got := removeElement(tt.slice, tt.index)
			if !tt.shouldPanic && !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("removeElement(%v, %d) = %v, want %v", tt.slice, tt.index, got, tt.expect)
			}
		})
	}
}
