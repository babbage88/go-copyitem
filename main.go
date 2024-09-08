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
	fmt.Printf("Destination Path: %s\n", dstfileinfo.path)

	srcfileinfo.GetFileInfo()
	dstfileinfo.GetFileInfo()

	filecopyjob.SourceFile = srcfileinfo
	filecopyjob.DestinationFile = dstfileinfo

	sizehumanread := filecopyjob.SourceFile.GetSizeInMB()
	dstsize := filecopyjob.DestinationFile.GetSizeInMB()

	fmt.Printf("size of %s is %.2f\n", filecopyjob.PrettyPrintSrc(), sizehumanread)
	fmt.Printf("Destination file %s size is %.2f\n", filecopyjob.PrettyPrintDst(), dstsize)

	filecopyjob.Start()

}
