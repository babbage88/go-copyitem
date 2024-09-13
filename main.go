package main

import (
	"flag"
)

func main() {
	source := flag.String("source", ".", "Source file or directory to copy")
	destination := flag.String("destination", "C:\temp", "Destination to Copy to.")
	widthProgress := flag.Int("width", 75, "Width of the Progress bar that get displayed in console.")
	flag.Parse()

	progressBarConfig := NewProgressBarConfig(WithProgressBarWidth(*widthProgress))

	filecopyjob := NewFileCopyJob(WithSourceFilePath(*source), WithDestinationFilePath(*destination), WithProgressBarConfig(progressBarConfig))

	filecopyjob.Start()

}
