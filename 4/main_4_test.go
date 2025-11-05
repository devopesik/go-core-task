package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name string
		a    []string
		b    []string
		want []string
	}{
		{
			name: "обычный случай",
			a:    []string{"apple", "banana", "cherry", "date"},
			b:    []string{"banana", "date"},
			want: []string{"apple", "cherry"},
		},
		{
			name: "сохранение дубликатов из a",
			a:    []string{"apple", "banana", "apple", "cherry", "banana"},
			b:    []string{"banana"},
			want: []string{"apple", "apple", "cherry"},
		},
		{
			name: "b содержит элементы, отсутствующие в a",
			a:    []string{"x", "y"},
			b:    []string{"z", "w", "x"},
			want: []string{"y"},
		},
		{
			name: "a пустой",
			a:    []string{},
			b:    []string{"banana", "date"},
			want: []string{},
		},
		{
			name: "b пустой",
			a:    []string{"apple", "banana"},
			b:    []string{},
			want: []string{"apple", "banana"},
		},
		{
			name: "оба пустые",
			a:    []string{},
			b:    []string{},
			want: []string{},
		},
		{
			name: "все элементы a есть в b",
			a:    []string{"a", "b", "c"},
			b:    []string{"c", "a", "b", "d"},
			want: []string{},
		},
		{
			name: "дубликаты в b (не влияют)",
			a:    []string{"x", "y", "z"},
			b:    []string{"x", "x", "x"},
			want: []string{"y", "z"},
		},
		{
			name: "строки с цифрами и спецсимволами",
			a:    []string{"43", "lead", "gno1", "test@test.com", ""},
			b:    []string{"lead", ""},
			want: []string{"43", "gno1", "test@test.com"},
		},
		{
			name: "чувствительность к регистру",
			a:    []string{"Apple", "apple"},
			b:    []string{"apple"},
			want: []string{"Apple"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := difference(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("difference(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
