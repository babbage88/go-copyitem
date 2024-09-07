package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

type FileCopyJob struct {
	SourceFile      FileInfoExtended `json:"sourcefileinfo"`
	DestinationFile FileInfoExtended `json:"sourcefileinfo"`
}

type FileInfoExtended struct {
	FsFileInfo fs.FileInfo `json:"fsfileinfo"`
	path       string      `json:"path"`
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

	src, err := os.Stat(*source)
	if err != nil {
		fmt.Errorf("Error when trying to Stat source file %s\n", src)
	}

	dst, err := os.Stat(*destination)
	if err != nil {
		fmt.Errorf("Error when trying to Stat source file %s\n", dst)
	}

	var srcfileinfo FileInfoExtended
	var dstfileinfo FileInfoExtended
	var filecopyjob FileCopyJob

	srcfileinfo.path = *source
	dstfileinfo.path = *destination
	srcfileinfo.FsFileInfo = src
	dstfileinfo.FsFileInfo = dst
	filecopyjob.SourceFile = srcfileinfo
	filecopyjob.DestinationFile = dstfileinfo

	sizehumanread := filecopyjob.SourceFile.GetSizeInMB()
	isSrcDir := filecopyjob.SourceFile.FsFileInfo.IsDir()
	isDstDir := filecopyjob.DestinationFile.FsFileInfo.IsDir()

	fmt.Printf("sizemb of %s is %v\n", filecopyjob.PrettyPrintSrc(), sizehumanread)

	if isSrcDir {
		fmt.Printf("The source file specified: %s is a Directory\n", filecopyjob.PrettyPrintSrc())
	} else {
		fmt.Printf("The source file specified: %s is not a Directory.\n", filecopyjob.PrettyPrintSrc())
	}

	if isDstDir {
		fmt.Printf("The destination file specified: %s is a Directory\n", filecopyjob.PrettyPrintDst())
	} else {
		fmt.Printf("The destination file specified: %s is not a Directory.\n", filecopyjob.PrettyPrintDst())
	}

}
