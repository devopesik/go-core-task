package main

import (
	"crypto/sha256"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			name:     "mixed types",
			input:    []any{100, 0144, 100.0, 0x64, true, complex(1, 2), "Hello"},
			expected: "100100100100true(1+2i)Hello",
		},
		{
			name:     "empty",
			input:    []any{},
			expected: "",
		},
		{
			name:     "only strings",
			input:    []any{"foo", "bar"},
			expected: "foobar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := merge(tt.input)
			if result != tt.expected {
				t.Errorf("merge() = %q, expected %q", result, tt.expected)
			}
		})
	}
}

func TestHash(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		expected [32]byte
	}{
		{
			name:  "known input",
			input: []rune("100100100100true(1+2i)Hello"),
			expected: func() [32]byte {
				salt := []byte("go-2024")
				data := append([]byte("100100100100true(1+2i)Hello"), salt...)
				return sha256.Sum256(data)
			}(),
		},
		{
			name:     "empty input",
			input:    []rune{},
			expected: sha256.Sum256([]byte("go-2024")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hash(tt.input)
			if result != tt.expected {
				t.Errorf("hash() = %x, expected %x", result, tt.expected)
			}
		})
	}
}
