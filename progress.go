package main

import (
	"fmt"
	"strings"
	"time"
)

type ProgressBarStats struct {
	PercentComplete      int
	Speed                float64
	EstimatedTimeRemaing time.Duration
	ETA                  time.Time
	Width                int
}

type ProgressBar interface {
	DrawProgressBar(barWidth int)
}

func (f *FileCopyJob) DrawProgressBar(barWidth int) {
	if f.ProgressCompleted < 0 {
		f.ProgressCompleted = 0
	} else if f.ProgressCompleted > 100 {
		f.ProgressCompleted = 100
	}

	// Calculate the number of "#" and "-" symbols to display
	filledBars := int(f.ProgressCompleted * float64(barWidth) / 100.0)
	emptyBars := barWidth - filledBars

	// Print the progress bar in place
	fmt.Printf("\r[%-*s] %.2f%%", barWidth, strings.Repeat("#", filledBars)+strings.Repeat("-", emptyBars), f.ProgressCompleted)

	// Move the cursor down one line, print speed, then move cursor back up
	fmt.Printf("\nSpeed: %s", f.PrettyPrintSpeedMB())
	fmt.Printf("\033[1A")
}

func DrawColoredString(s string, color int) string {
	coloredString := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, s)
	return coloredString
}

func DrawProgressBar(percentage float64, barWidth int, speed string) {
	if percentage < 0 {
		percentage = 0
	} else if percentage > 100 {
		percentage = 100
	}

	// Calculate the number of "#" and "-" symbols to display
	filledBars := int(percentage * float64(barWidth) / 100.0)
	emptyBars := barWidth - filledBars

	// Print the progress bar in place
	fmt.Printf("\r[%-*s] %.2f%%", barWidth, strings.Repeat("#", filledBars)+strings.Repeat("-", emptyBars), percentage)

	// Move the cursor down one line, print speed, then move cursor back up
	fmt.Printf("\nSpeed: %s", speed)
	fmt.Printf("\033[1A") // Move cursor up one line to overwrite the speed
}
