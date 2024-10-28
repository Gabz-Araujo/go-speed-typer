package stats

import (
	"time"
)

type GameStats struct {
	startTime    time.Time
	endTime      time.Time
	WPM          float64
	Accuracy     float64
	totalTyped   int
	correctTyped int
}

func NewGameStats() *GameStats {
	return &GameStats{
		startTime: time.Now(),
	}
}

func (gs *GameStats) Update(userInput, targetText string) {
	gs.totalTyped = len(userInput)
	gs.correctTyped = 0
	for i := 0; i < len(userInput) && i < len(targetText); i++ {
		if userInput[i] == targetText[i] {
			gs.correctTyped++
		}
	}

	if gs.totalTyped > 0 {
		gs.Accuracy = float64(gs.correctTyped) / float64(len(userInput)) * 100
	}

	if len(userInput) == len(targetText) {
		gs.endTime = time.Now()
		duration := gs.endTime.Sub(gs.startTime).Minutes()
		gs.WPM = float64(gs.totalTyped) / 5 / duration
	}
}

func (gs *GameStats) Reset() {
	gs.WPM = 0
	gs.Accuracy = 0
	gs.startTime = time.Now()
	gs.endTime = time.Time{}
	gs.totalTyped = 0
	gs.correctTyped = 0
}
