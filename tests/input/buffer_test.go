package input_test

import (
	"testing"

	"github.com/GabzAraujo/go-speed-typer/internal/input"
)

func TestRingBuffer(t *testing.T) {
	rb := input.NewRingBuffer(5)

	t.Cleanup(func() {
		rb.Clear()
	})

	t.Run("Initial State", func(t *testing.T) {
		if !rb.IsEmpty() {
			t.Error("New Buffer should be empty")
		}
	})

	tests := []struct {
		name     string
		input    []rune
		expected []rune
		count    int
	}{
		{"Single element", []rune{'a'}, []rune{'a'}, 1},
		{"Multiple elements", []rune{'a', 'b', 'c'}, []rune{'a', 'b', 'c'}, 3},
		{"Overflow", []rune{'a', 'b', 'c', 'd', 'e', 'f'}, []rune{'b', 'c', 'd', 'e', 'f'}, 5},
		{"Empty input", make([]rune, 0), make([]rune, 0), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := input.NewRingBuffer(5)
			for _, r := range tt.input {
				rb.Add(r)
			}

			if rb.Count() != tt.count {
				t.Errorf("Expected count %d, got %d", tt.count, rb.Count())
			}

			elements := rb.Get()
			if !runeSliceEqual(elements, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, elements)
			}
		})
	}

	t.Run("Buffer should be full if add more elements than the size", func(t *testing.T) {
		rb.Add('a')
		rb.Add('b')
		rb.Add('c')
		rb.Add('d')
		rb.Add('e')
		rb.Add('f')

		if !rb.IsFull() {
			t.Error("Buffer should be full")
		}
	})

	t.Run("Should move head and tail when adding to a full buffer", func(t *testing.T) {
		rb.Add('a')
		rb.Add('b')
		rb.Add('c')
		rb.Add('d')
		rb.Add('e')
		rb.Add('f')

		elements := rb.Get()
		expected := []rune{'b', 'c', 'd', 'e', 'f'}

		for i, e := range expected {
			if elements[i] != e {
				t.Errorf("Expected %c at position %d, got %c", e, i, elements[i])
			}
		}
	})

	t.Run("Should remove the last element from the RingBuffer", func(t *testing.T) {
		rb.Add('a')
		rb.Add('b')
		rb.Add('c')
		rb.Add('d')
		rb.Add('e')
		rb.Add('f')

		rb.RemoveLast()

		elements := rb.Get()
		expected := []rune{'b', 'c', 'd', 'e'}

		for i, e := range expected {
			if elements[i] != e {
				t.Errorf("Expected %c at position %d, got %c", e, i, elements[i])
			}
		}

		if rb.Count() != 4 {
			t.Errorf("Expected Count to be 4, got %d", rb.Count())
		}
	})
}

func runeSliceEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
