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

type ProgressBarConfig struct {
	Width              int
	FillCharacter      string
	RemainingCharacter string
}

type ProgressBar interface {
	DrawProgressBar()
	DrawColoredString(s string, color int) string
}

type ProgressBarConfigOptions func(*ProgressBarConfig)

func WithProgressBarWidth(width int) ProgressBarConfigOptions {
	return func(p *ProgressBarConfig) {
		p.Width = width
	}
}

func WithProgressFillCharacter(s string) ProgressBarConfigOptions {
	return func(p *ProgressBarConfig) {
		p.FillCharacter = s
	}
}

func WithProgressRemaingCharacter(s string) ProgressBarConfigOptions {
	return func(p *ProgressBarConfig) {
		p.RemainingCharacter = s
	}
}

func (f *ProgressBarConfig) DrawColoredString(s string, color int) string {
	coloredString := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, s)
	return coloredString
}

func NewProgressBarConfig(opts ...ProgressBarConfigOptions) *ProgressBarConfig {
	const (
		width = 50
	)
	progBarConf := &ProgressBarConfig{Width: width}
	progBarConf.FillCharacter = progBarConf.DrawColoredString("#", 92)
	progBarConf.RemainingCharacter = progBarConf.DrawColoredString("-", 96)

	for _, opt := range opts {

		opt(progBarConf)

	}

	return progBarConf
}

func (f *FileCopyJob) DrawProgressBar() {
	if f.ProgressCompleted < 0 {
		f.ProgressCompleted = 0
	} else if f.ProgressCompleted > 100 {
		f.ProgressCompleted = 100
	}

	// Calculate the number of "#" and "-" symbols to display
	filledBars := int(f.ProgressCompleted * float64(f.ProgressBarConfig.Width) / 100.0)
	emptyBars := f.ProgressBarConfig.Width - filledBars

	fillChar := f.DrawColoredString("#", 92)
	pctRemaingChar := f.DrawColoredString("-", 96)

	// Clear the current line before drawing the progress bar
	fmt.Printf("\r\033[K")

	// Print the progress bar
	fmt.Printf("\r[%-*s] %.2f%%", f.ProgressBarConfig.Width, strings.Repeat(fillChar, filledBars)+strings.Repeat(pctRemaingChar, emptyBars), f.ProgressCompleted)

	// Clear the next line for the speed display
	fmt.Printf("\n\033[K")

	// Print the speed in MB/s
	fmt.Printf("%s", f.PrettyPrintSpeedMB())

	if f.ProgressCompleted < 100 {
		// , if copy still running, move cursor up to clear progress bar on the next draw
		fmt.Printf("\033[1A")
	} else {
		fmt.Printf("\nFile Copy has completed.\n")
	}

}

func (f *FileCopyJob) DrawColoredString(s string, color int) string {
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
