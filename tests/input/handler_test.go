package input_test

import (
	"testing"

	"github.com/GabzAraujo/go-speed-typer/internal/input"
	tea "github.com/charmbracelet/bubbletea"
)

func TestInputHandler(t *testing.T) {
	t.Run("Initial State", func(t *testing.T) {
		handler := input.NewHandler(10)
		if !handler.IsEmpty() {
			t.Error("New handler should have an empty buffer")
		}
	})

	t.Run("Handle regular input", func(t *testing.T) {
		handler := input.NewHandler(10)
		handler.HandleInput(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
		handler.HandleInput(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
		handler.HandleInput(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})

		buffer := handler.GetInputBuffer()
		expected := []rune{'j', 'j', 'j'}
		if string(buffer) != string(expected) {
			t.Errorf("Expected buffer to be %v, got %v", string(expected), string(buffer))
		}
	})

	t.Run("Handle Backspace", func(t *testing.T) {
		handler := input.NewHandler(10)
		handler.HandleInput(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j', 'j', 'j'}})
		handler.HandleInput(tea.KeyMsg{Type: tea.KeyBackspace})

		buffer := handler.GetInputBuffer()
		expected := []rune{'j', 'j'}
		if string(buffer) != string(expected) {
			t.Errorf("Expected buffer to be %v, got %v", string(expected), string(buffer))
		}
	})

	t.Run("Handle Enter input", func(t *testing.T) {
		handler := input.NewHandler(10)
		handler.HandleInput(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j', 'j', 'j'}})
		handler.HandleInput(tea.KeyMsg{Type: tea.KeyEnter})

		if !handler.IsEmpty() {
			t.Error("Buffer should be cleared after Enter key")
		}
	})
}
