package input

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Handler struct {
	buffer    *RingBuffer
	textInput textinput.Model
}

func NewHandler(bufferSize int) *Handler {
	ti := textinput.New()
	ti.Focus()
	ti.Prompt = ""
	ti.EchoMode = textinput.EchoNormal
	ti.CharLimit = bufferSize

	return &Handler{
		buffer:    NewRingBuffer(bufferSize),
		textInput: ti,
	}
}

func (h *Handler) HandleInput(msg tea.Msg) (string, bool, tea.Cmd) {
	var cmd tea.Cmd
	var shouldQuit bool

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			shouldQuit = true
		case tea.KeyEnter:
			h.buffer.Clear()
		case tea.KeyBackspace:
			if !h.buffer.IsEmpty() {
				h.buffer.RemoveLast()
				h.textInput.SetValue(string(h.buffer.Get()))
			}
		default:

			if msg.Runes != nil {
				for _, r := range msg.Runes {
					h.buffer.Add(r)
				}
			}
		}
	}

	h.textInput, cmd = h.textInput.Update(msg)
	return string(h.buffer.Get()), shouldQuit, cmd
}

func (h *Handler) GetInputBuffer() []rune {
	return h.buffer.Get()
}

func (h *Handler) IsEmpty() bool {
	return h.buffer.IsEmpty()
}

func (h *Handler) Clear() {
	h.buffer.Clear()
	h.textInput.SetValue("")
}

func (h *Handler) Focus() {
	h.textInput.Focus()
}

func (h *Handler) Blur() {
	h.textInput.Blur()
}
