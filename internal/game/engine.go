package game

import (
	"fmt"

	"github.com/GabzAraujo/go-speed-typer/internal/input"
	"github.com/GabzAraujo/go-speed-typer/internal/stats"
	"github.com/GabzAraujo/go-speed-typer/internal/text"
	"github.com/GabzAraujo/go-speed-typer/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type GameState int

const (
	Start GameState = iota
	Typing
	Finish
)

type GameEngine struct {
	inputHandler *input.Handler
	textGen      *text.Generator
	stats        *stats.GameStats
	renderer     *ui.Renderer
	currentText  string
	userInput    string
	state        GameState
	cursorPos    int
}

func NewGameEngine() *GameEngine {
	return &GameEngine{
		inputHandler: input.NewHandler(100),
		textGen:      text.NewGenerator(),
		stats:        stats.NewGameStats(),
		state:        Start,
		renderer:     ui.NewRenderer(),
		currentText:  "",
		userInput:    "",
		cursorPos:    0,
	}
}

func (ge *GameEngine) Init() tea.Cmd {
	ge.currentText = ge.textGen.GenerateText()
	return nil
}

func (ge *GameEngine) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch ge.state {
	case Start:
		ge.UpdateStartScreen(msg)
	case Typing:
		ge.UpdateTypingScreen(msg)
	case Finish:
		ge.UpdateFinishScreen(msg)
	}

	return ge, nil
}

func (ge *GameEngine) UpdateStartScreen(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			ge.state = Typing
			ge.stats.Reset()
			ge.textGen.GenerateText()
			return ge, nil
		case tea.KeyCtrlC, tea.KeyEsc:
			return ge, tea.Quit
		}
	}
	return ge, nil
}

func (ge *GameEngine) UpdateTypingScreen(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var shouldQuit bool

	ge.userInput, shouldQuit, cmd = ge.inputHandler.HandleInput(msg)

	if shouldQuit {
		tea.Quit()
		return ge, tea.Quit
	}

	ge.cursorPos = len(ge.userInput)
	ge.stats.Update(ge.userInput, ge.currentText)

	if len(ge.userInput) == len(ge.currentText) {
		ge.inputHandler.Clear()
		ge.state = Finish
	}

	return ge, cmd
}

func (ge *GameEngine) UpdateFinishScreen(msg tea.Msg) (tea.Model, tea.Cmd) {
	var shouldQuit bool
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyShiftTab:
			ge.state = Start
			ge.currentText = ge.textGen.GenerateText()
		case tea.KeyCtrlR:
			ge.state = Typing
		case tea.KeyCtrlC, tea.KeyEsc:
			return ge, tea.Quit
		}
	}

	ge.inputHandler.Clear()
	ge.stats.Reset()
	ge.userInput = ""
	ge.cursorPos = 0
	if shouldQuit {
		return ge, tea.Quit
	}
	return ge, nil
}

func (ge *GameEngine) View() string {
	switch ge.state {
	case Start:
		return ge.renderer.RenderStartScreen()
	case Typing:
		return ge.renderer.RenderTypingScreen(
			ge.currentText,
			ge.userInput,
			ge.cursorPos,
			ge.stats,
		)
	case Finish:
		return ge.renderer.RenderFinishScreen(ge.stats)
	default:
		panic(fmt.Sprintf("unexpected game.GameState: %#v", ge.state))
	}
}
