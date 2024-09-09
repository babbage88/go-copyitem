package main

import (
	"flag"
	"fmt"
)

func main() {
	source := flag.String("source", ".", "Source file or directory to copy")
	destination := flag.String("destination", "C:\temp", "Destination to Copy to.")
	flag.Parse()

	var srcfileinfo FileInfoExtended
	var dstfileinfo FileInfoExtended
	var filecopyjob FileCopyJob

	srcfileinfo.path = *source
	dstfileinfo.path = *destination

	srcfileinfo.GetFileInfo()
	dstfileinfo.GetFileInfo()

	filecopyjob.SourceFile = srcfileinfo
	filecopyjob.DestinationFile = dstfileinfo

	sizehumanread := filecopyjob.SourceFile.GetSizeInMB()
	dstsize := filecopyjob.DestinationFile.GetSizeInMB()

	fmt.Printf("Source File size %s is %.2f MB\n", filecopyjob.PrettyPrintSrc(), sizehumanread)
	fmt.Printf("Destination file %s size is %.2f MB\n\n", filecopyjob.PrettyPrintDst(), dstsize)

	filecopyjob.Start()

}
