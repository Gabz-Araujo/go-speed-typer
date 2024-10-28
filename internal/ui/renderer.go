package ui

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/GabzAraujo/go-speed-typer/internal/stats"
	"github.com/charmbracelet/lipgloss"
)

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RenderStartScreen() string {
	var output strings.Builder
	output.WriteString(lipgloss.NewStyle().Bold(true).Render("Welcome to Speed Typer!\n\n"))
	output.WriteString("Press Enter to start typing\n")
	output.WriteString("Press Q to quit\n")
	return output.String()
}

func (r *Renderer) RenderFinishScreen(stats *stats.GameStats) string {
	var output strings.Builder
	output.WriteString(lipgloss.NewStyle().Bold(true).Render("Typing Session Complete!\n\n"))
	output.WriteString(fmt.Sprintf("WPM: %.2f\n", stats.WPM))
	output.WriteString(fmt.Sprintf("Accuracy: %.2f%%\n\n", stats.Accuracy))
	output.WriteString("Press Shift+Tab for a new session\n")
	output.WriteString("Press Ctrl+R to replay the same text\n")
	output.WriteString("Press Ctrl+Q to quit\n")
	return output.String()
}

func (r *Renderer) renderTypingArea(
	ghostText, userInput string,
	cursorPos int,
) string {
	var output strings.Builder

	for i, ghostChar := range ghostText {
		if i < cursorPos {
			userChar, _ := utf8.DecodeRuneInString(userInput[i:])
			if userChar == ghostChar {
				output.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render(string(userChar)))
			} else {
				output.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(string(userChar)))
			}
		} else if i == cursorPos {
			output.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Render(string(ghostChar)))
		} else {
			output.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(string(ghostChar)))
		}
	}

	return output.String()
}

func (r *Renderer) RenderTypingScreen(
	ghostText, userInput string,
	cursorPos int,
	stats *stats.GameStats,
) string {
	var output strings.Builder

	frame := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		Padding(5).
		Align(lipgloss.Center)

	typingArea := r.renderTypingArea(ghostText, userInput, cursorPos)

	statsArea := fmt.Sprintf("WPM: %.2f | Accuracy: %.2f%%", stats.WPM, stats.Accuracy)

	content := lipgloss.JoinVertical(lipgloss.Center,
		typingArea,
		statsArea,
	)

	output.WriteString(frame.Render(content))
	return output.String()
}
