package main

import (
	"reflect"
	"testing"
)

func TestIntercept(t *testing.T) {
	tests := []struct {
		name      string
		a         []int
		b         []int
		wantHas   bool
		wantSlice []int
	}{
		{
			name:      "обычное пересечение",
			a:         []int{65, 3, 58, 678, 64},
			b:         []int{64, 2, 3, 43},
			wantHas:   true,
			wantSlice: []int{64, 3},
		},
		{
			name:      "дубликаты в a — результат уникален",
			a:         []int{1, 1, 2, 2, 3},
			b:         []int{2, 3, 4},
			wantHas:   true,
			wantSlice: []int{2, 3},
		},
		{
			name:      "дубликаты в b — результат уникален",
			a:         []int{2, 3, 5},
			b:         []int{2, 2, 3, 3, 4},
			wantHas:   true,
			wantSlice: []int{2, 3},
		},
		{
			name:      "дубликаты в обоих — результат уникален",
			a:         []int{1, 1, 2, 2},
			b:         []int{2, 2, 1, 1},
			wantHas:   true,
			wantSlice: []int{2, 1}, // порядок из b
		},
		{
			name:      "нет пересечения",
			a:         []int{1, 2, 3},
			b:         []int{4, 5, 6},
			wantHas:   false,
			wantSlice: []int{},
		},
		{
			name:      "a пустой",
			a:         []int{},
			b:         []int{1, 2, 3},
			wantHas:   false,
			wantSlice: []int{},
		},
		{
			name:      "b пустой",
			a:         []int{1, 2, 3},
			b:         []int{},
			wantHas:   false,
			wantSlice: []int{},
		},
		{
			name:      "оба пустые",
			a:         []int{},
			b:         []int{},
			wantHas:   false,
			wantSlice: []int{},
		},
		{
			name:      "один общий элемент",
			a:         []int{42},
			b:         []int{10, 20, 42, 30},
			wantHas:   true,
			wantSlice: []int{42},
		},
		{
			name:      "все элементы совпадают",
			a:         []int{1, 2, 3},
			b:         []int{3, 2, 1},
			wantHas:   true,
			wantSlice: []int{3, 2, 1}, // порядок из b
		},
		{
			name:      "отрицательные числа",
			a:         []int{-1, -2, 0, 5},
			b:         []int{0, -2, 10},
			wantHas:   true,
			wantSlice: []int{0, -2},
		},
		{
			name:      "нули",
			a:         []int{0, 0, 1},
			b:         []int{0, 2, 0},
			wantHas:   true,
			wantSlice: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHas, gotSlice := intercept(tt.a, tt.b)

			if gotHas != tt.wantHas {
				t.Errorf("intercept(%v, %v): has = %v, want %v", tt.a, tt.b, gotHas, tt.wantHas)
			}

			// В Go nil-слайс и пустой слайс считаются равными через reflect.DeepEqual
			if !reflect.DeepEqual(gotSlice, tt.wantSlice) {
				t.Errorf("intercept(%v, %v): slice = %v, want %v", tt.a, tt.b, gotSlice, tt.wantSlice)
			}
		})
	}
}
