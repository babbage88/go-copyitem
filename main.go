package main

import (
	"flag"
	"fmt"
)

type FileCopyJob struct {
	SourceFile      FileInfoExtended `json:"sourcefileinfo"`
	DestinationFile FileInfoExtended `json:"sourcefileinfo"`
}

type ColorizedSrcPath interface {
	PrettyPrintSrc() string
}

type ColorizedDstPath interface {
	PrettyPrintDst() string
}

func (f *FileCopyJob) PrettyPrintSrc() string {
	coloredsource := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 96, f.SourceFile.path)
	return coloredsource
}

func (f *FileCopyJob) PrettyPrintDst() string {
	colordestination := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 92, f.DestinationFile.path)
	return colordestination
}

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

	fmt.Printf("sizemb of %s is %.2f\n", filecopyjob.PrettyPrintSrc(), sizehumanread)
	fmt.Printf("Destination file %s size is %.2f\n", filecopyjob.PrettyPrintDst(), dstsize)

	//filecopyjob.GetCopyProgressPercentStr()

}
