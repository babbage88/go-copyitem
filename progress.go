package main

import (
	"fmt"
	"strings"
)

// DrawProgressBar draws the progress bar in the console.
func DrawProgressBar(percentage float64, barWidth int) {
	// Clamp the percentage value between 0 and 100
	if percentage < 0 {
		percentage = 0
	} else if percentage > 100 {
		percentage = 100
	}

	// Calculate the number of "#" and "-" symbols to display
	filledBars := int(percentage * float64(barWidth) / 100.0)
	emptyBars := barWidth - filledBars

	// Build the progress bar string
	progressBar := fmt.Sprintf("\r[%-*s] %.2f%%", barWidth, strings.Repeat("#", filledBars)+strings.Repeat("-", emptyBars), percentage)

	// Print the progress bar in place
	fmt.Print(progressBar)
}
