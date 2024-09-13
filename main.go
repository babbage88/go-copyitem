package main

import (
	"flag"
	"fmt"
)

func main() {
	source := flag.String("source", ".", "Source file or directory to copy")
	destination := flag.String("destination", "C:\temp", "Destination to Copy to.")
	widthProgress := flag.Int("width", 50, "Width of the Progress bar that get displayed in console.")
	flag.Parse()

	progressBarConfig := NewProgressBarConfig(WithProgressBarWidth(*widthProgress))

	filecopyjob := NewFileCopyJob(WithSourceFilePath(*source), WithDestinationFilePath(*destination), WithProgressBarConfig(progressBarConfig))

	sizehumanread := filecopyjob.SourceFile.GetSizeInMB()
	dstsize := filecopyjob.DestinationFile.GetSizeInMB()

	fmt.Printf("Source File %s size is %.2f MB\n", filecopyjob.PrettyPrintSrc(), sizehumanread)
	fmt.Printf("Destination file %s size is %.2f MB\n\n", filecopyjob.PrettyPrintDst(), dstsize)

	filecopyjob.Start()

}
